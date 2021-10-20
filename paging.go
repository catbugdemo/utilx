package main

import (
	"strconv"
	"strings"
)

// ToLimitOffset 用于返回相应 limit , offset
func ToLimitOffset(sizeStr string, pageStr string, count int) (limit int, offset int) {
	size, _ := strconv.Atoi(sizeStr)
	page, _ := strconv.Atoi(pageStr)

	if size == 0 {
		size = 10
	}

	if page == 0 {
		page = 1
	}

	if count == 0 {
		return size, 0
	}

	var pageMax int
	if count%size == 0 {
		pageMax = count / size
	} else {
		pageMax = count/size + 1
	}

	if pageMax <= size {
		page = pageMax
	}

	offset = size * (page - 1)
	return size, offset
}

//  ParseOrderBy 解析为相应的
//  order_by=-id,create_at
func ParseOrderBy(orderBy string) string {
	tmp := make([]string, 0, 2)
	orders := strings.Split(orderBy, ",")
	for _, order := range orders {
		order = strings.TrimSpace(order)
		if strings.HasPrefix(order, "-") {
			tmp = append(tmp, order[1:]+" DESC")
		} else {
			tmp = append(tmp, order[1:]+" ASC")
		}
	}
	return strings.Join(tmp, ",") + " "
}
