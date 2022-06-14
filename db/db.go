/*
	The main purpose of this package is to provide db access for all other packages. Unified db client and stuff..
*/

package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	uri = ""
)

// Client has Mutex to provide access for multiple goroutines
type Client struct {
	Mutex  sync.Mutex
	Client *mongo.Client
}

// Init sets default URI for mongo
func Init(mongoURI string) {
	uri = mongoURI
}

// NewClient returns new client, connected to database
func NewClient() *Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Panicln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Panicln(err)
	}
	var dbc Client
	dbc.Mutex.Lock()
	dbc.Client = client
	dbc.Mutex.Unlock()
	return &dbc
}

// InsertData just inserts one value
func (d *Client) InsertData(data interface{}, db, collection string) error {
	d.Mutex.Lock()
	_, err := d.Client.Database(db).Collection(collection).InsertOne(context.Background(), data)
	d.Mutex.Unlock()
	return err
}

// InsertMany inserts multiple values
func (d *Client) InsertMany(data []interface{}, db, collection string) error {
	d.Mutex.Lock()
	_, err := d.Client.Database(db).Collection(collection).InsertMany(context.Background(), data)
	d.Mutex.Unlock()
	return err
}

// FindData finds data and decodes to the given pointer
func (d *Client) FindData(filter interface{}, toDecode interface{}, db, collection string) error {
	d.Mutex.Lock()
	err := d.Client.Database(db).Collection(collection).FindOne(context.Background(), filter).Decode(toDecode)
	d.Mutex.Unlock()
	return err
}

// DeleteItem deletes one item
func (d *Client) DeleteItem(filter interface{}, db, collection string) error {
	d.Mutex.Lock()
	_, err := d.Client.Database(db).Collection(collection).DeleteOne(context.Background(), filter)
	d.Mutex.Unlock()
	return err
}

// DeleteItems deletes many items
func (d *Client) DeleteItems(filter interface{}, db, collection string) error {
	d.Mutex.Lock()
	_, err := d.Client.Database(db).Collection(collection).DeleteMany(context.Background(), filter)
	d.Mutex.Unlock()
	return err
}

// UpdateData just updates one item by given filter
func (d *Client) UpdateData(filter interface{}, update interface{}, db, collection string) error {
	d.Mutex.Lock()
	_, err := d.Client.Database(db).Collection(collection).UpdateOne(context.Background(), filter, update)
	d.Mutex.Unlock()
	return err
}

// FindMany finds many entries
func (d *Client) FindMany(filter interface{}, db, collection string) (*mongo.Cursor, error) {
	d.Mutex.Lock()
	cur, err := d.Client.Database(db).Collection(collection).Find(context.Background(), filter)
	d.Mutex.Unlock()
	return cur, err
}
