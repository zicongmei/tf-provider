package provider

import (
	"fmt"
	"log"
	"time"

	"git.mills.io/prologic/bitcask"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getDBFile(d *schema.ResourceData, m any, item string) string {
	path := m.(*DBClient).Path1
	if item == resource2 {
		path = m.(*DBClient).Path2
	}
	log.Println("Using DB path %q.", path)
	return path
	//return d.Get("db_file").(string)
	//return "/tmp/tf_demo_db"
}

func resource1Create(d *schema.ResourceData, m any) error {
	return resourceCreate(d, m, resource1)
}
func resource2Create(d *schema.ResourceData, m any) error {
	return resourceCreate(d, m, resource2)
}
func resourceCreate(d *schema.ResourceData, m any, item string) error {
	dbFile := getDBFile(d, m, item)
	db, err := bitcask.Open(dbFile)
	if err != nil {
		return fmt.Errorf("failed to open db file %q: %v", dbFile, err)
	}
	defer db.Close()
	key := d.Get(resourceKey).(string)
	value := d.Get(resourceValue).(string)
	err = db.Put([]byte(key), []byte(value))
	if err != nil {
		return fmt.Errorf("failed to write db key %q: %v", key, err)
	}

	log.Printf("Created key %q.", d.Get(resourceKey).(string))
	timeNow := time.Now().Format(time.DateTime)
	d.SetId(uuid.New().String())
	d.Set("last_updated", timeNow)
	d.Set("created", timeNow)

	return nil
}

func resource1Update(d *schema.ResourceData, m any) error {
	return resourceUpdate(d, m, resource1)
}
func resource2Update(d *schema.ResourceData, m any) error {
	return resourceUpdate(d, m, resource2)
}
func resourceUpdate(d *schema.ResourceData, m any, item string) error {
	dbFile := getDBFile(d, m, item)
	db, err := bitcask.Open(dbFile)
	if err != nil {
		return fmt.Errorf("failed to open db file %q: %v", dbFile, err)
	}
	defer db.Close()
	key := d.Get(resourceKey).(string)
	value := d.Get(resourceValue).(string)
	err = db.Put([]byte(key), []byte(value))
	if err != nil {
		return fmt.Errorf("failed to write db key %q: %v", key, err)
	}

	log.Printf("Updated key %q.", key)
	timeNow := time.Now().Format(time.DateTime)
	d.Set("last_updated", timeNow)

	return nil
}

func resource1Read(d *schema.ResourceData, m any) error {
	return resourceRead(d, m, resource1)
}
func resource2Read(d *schema.ResourceData, m any) error {
	return resourceRead(d, m, resource2)
}
func resourceRead(d *schema.ResourceData, m any, item string) error {
	dbFile := getDBFile(d, m, item)
	db, err := bitcask.Open(dbFile)
	if err != nil {
		return fmt.Errorf("failed to open db file %q: %v", dbFile, err)
	}
	defer db.Close()
	key := d.Get(resourceKey).(string)
	out, err := db.Get([]byte(key))
	if err != nil {
		return fmt.Errorf("failed to read db key %q: %v", key, err)
	}

	log.Print("Got value %q for key %q", string(out), key)

	d.Set("last_updated", time.Now().Format(time.DateTime))

	return nil
}
func resource1Delete(d *schema.ResourceData, m any) error {
	return resourceDelete(d, m, resource1)
}
func resource2Delete(d *schema.ResourceData, m any) error {
	return resourceDelete(d, m, resource2)
}
func resourceDelete(d *schema.ResourceData, m any, item string) error {
	dbFile := getDBFile(d, m, item)
	db, err := bitcask.Open(dbFile)
	if err != nil {
		return fmt.Errorf("failed to open db file %q: %v", dbFile, err)
	}
	defer db.Close()

	key := d.Get(resourceKey).(string)
	err = db.Delete([]byte(key))
	if err != nil {
		return fmt.Errorf("failed to delete db key %q: %v", key, err)
	}
	log.Printf("Deleted key %q.", key)

	d.SetId("")
	return nil
}
