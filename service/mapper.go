package service

func ToEntity(dto OrderDto) Order {
	return Order{
		dto.Email,
		dto.Room,
		dto.From,
		dto.To,
	}
}

func ToDto(entity Order) OrderDto {
	return OrderDto{
		entity.userEmail,
		entity.room,
		entity.from,
		entity.to,
	}
}

func ToDtos(entities []Order) []OrderDto {
	dtos := []OrderDto{}
	for _, item := range entities {
		dtos = append(dtos, ToDto(item))
	}
	return dtos
}
