package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Функція представляє процес курця
func smoker(c0 int, list *list.List, wg *sync.WaitGroup, l *sync.Mutex) {
	for {
		if list.Len() != 0 {
			l.Lock()
			if list.Len() == 2 {
				c1 := list.Front().Value
				list.Remove(list.Front())

				c2 := list.Front().Value
				list.Remove(list.Front())

				if c1 != c0 && c2 != c0 {
					// Курець може скрутити і курити цигарку, якщо він має два компоненти, які не співпадають з його власним компонентом
					fmt.Printf("Курець номер -%d скрутив сигарету і скурив її \n", c0+1)
					time.Sleep(1000 * time.Millisecond)
					wg.Done() // Завершуємо курця, після того як він закінчив курити.
				} else {
					// Якщо компоненти не підходять для куріння даного курця, то повертаємо їх на стіл
					fmt.Printf(" Курець номер-%d не зміг покурити \n", c0+1)
					list.PushBack(c1)
					list.PushBack(c2)
				}
			}
			l.Unlock()
		}
	}
}

// Функція представляє процес посередника, який кладає компоненти на стіл
func producer(ls *list.List, wg *sync.WaitGroup) {
	for {
		v0 := rand.Intn(3)
		v1 := rand.Intn(3)

		if v0 == v1 {
			v1 = (v1 + 1) % 3
		}

		ls.PushBack(v0)
		ls.PushBack(v1)

		fmt.Printf("Поклали предмети %d %d \n", v0, v1)

		wg.Add(1) // Збільшуємо лічильник очікування для курця
		wg.Wait() // Очікуємо, поки курець закінчить курити, перш ніж додати ще компоненти на стіл
	}
}

func main() {
	var list list.List

	var wg sync.WaitGroup
	//Використовуємо Mutex який контролює доступ до спільних даних
	var l sync.Mutex

	// Запускаємо горутини для трьох курців та посередника
	go smoker(0, &list, &wg, &l)
	go smoker(1, &list, &wg, &l)
	go smoker(2, &list, &wg, &l)

	go producer(&list, &wg)

	for {
		// Головний цикл main
	}
}
