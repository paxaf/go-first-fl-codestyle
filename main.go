package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func attack(name, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randInt(3, 5))
	}

	if class == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randInt(5, 10))
	}

	if class == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", name, 5+randInt(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func defence(name, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randInt(5, 10))
	} else if class == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randInt(-2, 2))
	} else if class == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", name, 10+randInt(2, 5))
	}
	return "неизвестный класс персонажа"
}

// обратите внимание на "if else" и на "else"
func special(name, class string) string {
	if class == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", name, 80+25)
	} else if class == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", name, 5+40)
	} else if class == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", name, 10+30)
	}
	return "неизвестный класс персонажа"
}

// здесь обратите внимание на имена параметров
func trainingStart(name, class string) string {
	if class == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", name)
	}

	if class == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", name)
	}

	if class == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", name)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)
		switch {
		case cmd == "attack":
			fmt.Println(attack(name, class))
		case cmd == "defence":
			fmt.Println(defence(name, class))
		case cmd == "special":
			fmt.Println(special(name, class))
		default:
			fmt.Println("неизвестная команда")
		}
	}
	return "тренировка окончена"
}

// обратите внимание на имя функции и имена переменных
func classChoice() string {
	var approveChoice string
	var class string

	for approveChoice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &class)
		if class == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if class == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if class == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoice)
		approveChoice = strings.ToLower(approveChoice)
	}
	return class
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var name string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Здравствуй, %s\n", name)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	class := classChoice()

	fmt.Println(trainingStart(name, class))
}
