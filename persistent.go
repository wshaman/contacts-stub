package stub_contacts

import "log"

var persistent YellowPages

func LoaPersistent() YellowPages {
	var err error
	if persistent == nil {
		if persistent, err = NewYellowPages(); err != nil {
			log.Fatal(err)
		}
		if err = Populate(persistent); err != nil {
			log.Fatal(err)
		}
	}
	return persistent

}
