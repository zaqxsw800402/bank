package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"red/dto"
	"sort"
	"strconv"
	"time"
)

var ctx = context.Background()

//var RC *redis.Client

type Database struct {
	RC *redis.Client
}

func New(client *redis.Client) Database {
	return Database{client}
}

func GetClient(host string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: "",
		DB:       0,
	})

	Pong, err := client.Ping(ctx).Result()
	log.Println("redis connection " + Pong)
	return client, err
}

type Customer struct {
	Id          uint   `redis:"Id" json:"customer_id"`
	Name        string `redis:"Name" json:"name"`
	City        string `redis:"City" json:"city"`
	DateOfBirth string `redis:"DateOfBirth" json:"date_of_birth"`
	Status      string `redis:"Status" json:"status"`
}

func (d Database) getUserValueForCustomer(customerID string) string {
	return fmt.Sprintf("customer:%s", customerID)
}

func (d Database) getUserKeyForCustomer(userID int) string {
	return fmt.Sprintf("%d:customer", userID)
}

// SaveCustomer 存資料到redis
func (d Database) SaveCustomer(ctx context.Context, userID int, c dto.CustomerResponse) {
	userKey := d.getUserKeyForCustomer(userID)
	// add customer to user id
	userValue := d.getUserValueForCustomer(strconv.Itoa(int(c.Id)))

	result := d.RC.Exists(ctx, userKey)
	d.RC.SAdd(ctx, userKey, userValue)
	if result.Val() == 0 {
		d.RC.Expire(ctx, userKey, time.Hour*1)
	}

	// 存進資料
	d.RC.HSet(ctx, userValue,
		"Id", c.Id,
		"Name", c.Name,
		"City", c.City,
		"DateOfBirth", c.DateOfBirth,
		"Status", c.Status)
	// 設定過期時間
	d.RC.Expire(ctx, userValue, time.Hour*1)
}

func (d Database) SaveAllCustomers(ctx context.Context, userid int, customers []dto.CustomerResponse) {
	for _, customer := range customers {
		d.SaveCustomer(ctx, userid, customer)
	}
}

func (d Database) GetAllCustomers(ctx context.Context, userid int) ([]dto.CustomerResponse, error) {
	userKey := d.getUserKeyForCustomer(userid)
	members := d.RC.SMembers(ctx, userKey)

	customers := make([]dto.CustomerResponse, 0)

	for _, member := range members.Val() {
		customer, err := d.GetCustomer(ctx, member)
		switch {
		case err != nil:
			return nil, err
		case customer == nil:
		case customer.Status == "":
		default:
			customers = append(customers, dto.CustomerResponse{
				Id:          customer.Id,
				Name:        customer.Name,
				City:        customer.City,
				DateOfBirth: customer.DateOfBirth,
				Status:      customer.Status,
			})
		}
		//if err != nil {
		//	return nil, err
		//}
		//customers = append(customers, dto.CustomerResponse{
		//	Id:          customer.Id,
		//	Name:        customer.Name,
		//	City:        customer.City,
		//	DateOfBirth: customer.DateOfBirth,
		//	Status:      customer.Status,
		//})
	}

	sort.SliceStable(customers, func(i int, j int) bool {
		return customers[i].Id < customers[j].Id
	})
	return customers, nil
}

// GetCustomer 讀取使用者資料
func (d Database) GetCustomer(ctx context.Context, customerKey string) (*Customer, error) {
	var customer Customer
	err := d.RC.HGetAll(ctx, customerKey).Scan(&customer)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &customer, nil
}

func (d Database) DeleteCustomer(customerID string) error {
	userValue := d.getUserValueForCustomer(customerID)
	result := d.RC.HSet(ctx, userValue, "Status", 0)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
