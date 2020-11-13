package dbTypes

type DynamoDBGetter interface {
	GetItem(partitionKey interface{}, sortKey interface{}) (interface{}, error)
}

type DynamoDBInserter interface {
	InsertItem(interface{}) error
}

type DynamoDBUpdater interface {
	UpdateItem(interface{}) error
}

type DynamoDBExistsChecker interface {
	CheckItemExists(partitionKey interface{}, sortKey interface{}) (bool, error)
}

type DynamoDBLister interface {
	ListItems(partitionKey interface{}, keyPrefix interface{}) (interface{}, error)
}
