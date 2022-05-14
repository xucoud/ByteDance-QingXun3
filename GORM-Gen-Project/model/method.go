package model

import "gorm.io/gen"

type TUserMethod interface {
	// sql(select * from users where uuid=@uuid )
	FindByUUID(uuid string) (gen.T, error)

	// sql(select * from users)
	FindAll() (gen.M, error)

	// sql(insert into users (uuid, name ,age, version) values (@uuid,@name,@age,@version) )
	InsertValue(uuid string, age int, name string, version int) error

	// sql( update users set name = @name,age = @age, version = @version where uuid = @uuid )
	UpdateVersion(uuid string, age int, name string, version int) error

	// sql(delete from users where uuid = @uuid )
	DeleteValue(uuid string) error

	// sql(select sum(version) from users where version = (select version from users group by version ORDER BY version desc limit 1) )
	FindMaxVersion() (int, error)
}

// sql(select * from users where version = ( select version from users having 1 order by version desc) )
