package primitives

type IDomainEventDispatcher interface {
	DispatchAndClearEvents(entities []BaseEntity) error
}
