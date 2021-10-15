package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Character struct {
	nom       string
	class     string
	pvmax     int
	pv        int
	inv       []string
	nbpot     int
	gold      int
	skill     []string
	equipment Equipment
	first     string
}

type Equipment struct {
	second string
	head   string
	chest  string
}

type Monstre struct {
	nom     string
	pv      int
	pvmax   int
	attaque int
}

func clearScreen() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func main() {
	var c1 Character
	c1.skill = append(c1.skill, "Coup de poing")
	clearScreen()

	sBar := "█"
	sVide := "                                                   "

	for i := 0; i <= 100; i++ {
		fmt.Printf("\n\n\n\n\n\n")
		fmt.Print("[|" + sBar + sVide + strconv.Itoa(i) + "% Complété.]")
		time.Sleep(30 * time.Millisecond)
		clearScreen()
		if i%2 == 0 {
			sBar += "█"
			sVide = sVide[:len(sVide)-1]
		}
	}
	clearScreen()
	fmt.Printf("\n\n\n\n\n\n")
	fmt.Print("[ Chargement du monde terminé. ]")
	time.Sleep(1000 * time.Millisecond)
	clearScreen()
	fmt.Print("\n\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("[ ✨ ~ Bienvenue dans le monde de Syneria ! ~ ✨ ] ")
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	time.Sleep(3 * time.Second)
	clearScreen()

	c1.pseudo()
}

func (c *Character) pseudo() {
START:
	fmt.Print("\n\n\n\n\n\n💬  ~ Quel est ton nom jeune aventurier : ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		clearScreen()
		fmt.Print("\n\n\n\n[ 💬  ~ Tu te nomme ", scanner.Text(), " c'est bien ça ? ]\n")
		fmt.Print("\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
		fmt.Print("Tapez [1] pour :  Oui, c'est bien mon nom.\n")
		fmt.Print("Tapez [2] pour :  Non, ce n'est pas mon nom\n")
		fmt.Print("\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		scanner1 := bufio.NewScanner(os.Stdin)
		for scanner1.Scan() {
			if scanner1.Text() == "1" {
				c.nom = scanner.Text()
				c.create()
			} else if scanner1.Text() == "2" {
				clearScreen()
				goto START
			} else {
				fmt.Println("Il faut que tu écrive 1 ou 2.")
			}
		}
	}
}

func (c Character) create() {
	clearScreen()
	fmt.Print("\n\n\n[ 💬  ~ Salutation ", c.nom, ", dis moi qui es-tu ? ] \n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Tapez [1] pour :  Je suis un Elfe.\n")
	fmt.Print("Tapez [2] pour :  Je suis un Humain.\n")
	fmt.Print("Tapez [3] pour :  Je suis un Nain.\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	c.gold = 150
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.class = "Elfe"
			c.pvmax = 80
			c.pv = 40
			c.first = "Epée Elfique"
			c.inv = []string{"Epée Elfique", " Potion de Soin", " Potion de Soin", " Potion de Soin"}
			clearScreen()
			fmt.Print(" \n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n 💬  ~ Tu es donc un Elfe, j'en étais sur !\n Tu possède alors 40 PV (Points de Vie) & 80 PV Max.\n Enfin, ton inventaire est constituté d'une Epée Elfique & 3 Potions de soin.\n\n Bonne chance à toi !\n\n (Tapez [1] pour accéder au Menu)\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
			c.gomenu()
		} else if scanner.Text() == "2" {
			c.class = "Humain"
			c.pvmax = 100
			c.pv = 50
			c.first = "Epée Classique"
			c.inv = []string{"Epée Classique", " Bouclier ", " Potion de Soin"}
			clearScreen()
			fmt.Print(" \n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n 💬  ~ Tu es donc un Humain, j'en étais sur !\n Tu possède alors 50 PV (Points de Vie) & 120 PV Max.\n Enfin, ton inventaire est constituté d'une Epée Classique, un Bouclier Solide & 1 Potion de Soin.\n\n Bonne chance à toi !\n\n (Tapez [1] pour accéder au Menu)\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
			c.gomenu()
		} else if scanner.Text() == "3" {
			c.class = "Nain"
			c.pvmax = 120
			c.pv = 60
			c.first = "Grosse Hâche"
			c.inv = []string{"Grosse Hâche", " Potion de Soin", " Potion de Soin"}
			clearScreen()
			fmt.Print(" \n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n 💬  ~ Tu es donc un Nain, j'en étais sur !\n Tu possède alors 60 PV (Points de Vie) & 120 PV Max.\n Enfin, ton inventaire est constituté d'une Grosse Hâche & 2 Potions de soin.\n\n Bonne chance à toi !\n\n (Tapez [1] pour accéder au Menu)\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
			c.gomenu()
		} else {
			fmt.Println("Il faut que tu écrive 1, 2 ou 3.")
		}
	}
}

func (c *Character) gomenu() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.menu()
		} else {
			fmt.Print("Il faut que tu écrive 1 pour accéder au Menu.")
		}
	}
}

func (c *Character) menu() {
	clearScreen()
	fmt.Print("\n\n📖  ~ Menu :")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Tapez [1] pour :  Afficher les informations du personnage\n")
	fmt.Print("Tapez [2] pour :  Accéder au contenu de l’inventaire\n")
	fmt.Print("Tapez [3] pour :  Accéder au Marchand du village\n")
	fmt.Print("Tapez [4] pour :  Accéder au Forgeron du village\n")
	fmt.Print("Tapez [5] pour :  Combattre\n\n")
	fmt.Print("Tapez [6] pour :  Quitter\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.displayInfo()
		} else if scanner.Text() == "2" {
			c.accesInventory()
		} else if scanner.Text() == "3" {
			c.merchant()
		} else if scanner.Text() == "4" {
			c.blacksmith()
		} else if scanner.Text() == "5" {
			c.combat()
		} else if scanner.Text() == "6" {
			c.end()
		} else {
			fmt.Println("Il faut que tu écrive 1, 2, 3, 4, 5 ou 6.")
		}
	}
	c.dead()
}

func (c Character) displayInfo() {
	clearScreen()
	fmt.Print("\n📜 ~ Informations du personnage :\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Nom : ", c.nom, "\n")
	fmt.Print("Classe : ", c.class, "\n")
	fmt.Print("PV : ", c.pv, "\n")
	fmt.Print("PV Max : ", c.pvmax, "\n")
	fmt.Print("Skill : ", c.skill, "\n")
	fmt.Print("\n(Tapez [1] pour retourner au Menu)")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.menu()
		} else {
			fmt.Print("Il faut que tu écrive 1 pour retourner au Menu.")
		}
	}
}

func (c *Character) nnbpot() int {
	var nb int
	for i := 0; i < len(c.inv); i++ {
		if c.inv[i] == "Potion de soin" {
			nb += 1
		}
	}
	return nb
}

func (c *Character) accesInventory() {
	clearScreen()
	fmt.Print("\n\n💼 ~ Inventaire du personnage :\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Items présents : ", c.inv, "\n")
	fmt.Print("\n(Tapez [1] pour retourner au Menu)")
	fmt.Print("\n(Tapez [2] pour utiliser une Potion de Soin)")
	fmt.Print("\n(Tapez [3] pour utiliser une Potion de Poison)")
	fmt.Print("\n(Tapez [4] pour utiliser le Livre de Sort Magique)")
	fmt.Print("\n(Tapez [5] pour accéder à l'Armurerie)")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.menu()
		} else if scanner.Text() == "2" {
			c.takePot()
		} else if scanner.Text() == "3" {
			c.takePotPoison()
		} else if scanner.Text() == "4" {
			c.spellBook()
		} else if scanner.Text() == "5" {
			c.armory()
		} else {
			fmt.Print("Il faut que tu écrive 1, 2, 3, 4 ou 5.")
		}
	}
	c.dead()
}
func (c *Character) takePotPoison() {
	clearScreen()
	fmt.Print("\n\n[ 💬  ~  Voulez vous vraiment utiliser une Potion de Poison, celle-ci vous infligera des dégâts irréverisibles, vous êtes sûr ? ~ ] ")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Tapez [1] pour : Oui \n")
	fmt.Print("Tapez [2] pour : Non \n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.takePotPoison2()
		} else if scanner.Text() == "2" {
			c.accesInventory()
		} else {
			fmt.Println("Il faut que tu écrive 1 ou 2.")
		}
	}
}

func (c *Character) takePotPoison2() {
	foundPotion2 := false
	for _, letters := range c.inv {
		if letters == " Potion de Poison" {
			foundPotion2 = true
			clearScreen()
			c.pv -= 10
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ☠️ ", "  ~ Aie... La Potion de Poison fait effet, vous perdez [10 PV] et êtes actuellement à [", c.pv, " PV] ! (Temps restant : 2s) ~ ☠️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(1 * time.Second)
			clearScreen()
			c.pv -= 10
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ☠️ ", "  ~ Aie... La Potion de Poison fait effet, vous perdez [10 PV] et êtes actuellement à [", c.pv, " PV] ! (Temps restant : 1s) ~ ☠️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(1 * time.Second)
			clearScreen()
			c.pv -= 10
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ☠️ ", "  ~ Aie... La Potion de Poison fait effet, vous perdez [10 PV] et êtes actuellement à [", c.pv, " PV] ! (Temps restant : 0s) ~ ☠️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(1 * time.Second)
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ❗❗", "~ Suite à cette Potion de Poison, vous passez maintenant à [", c.pv, " PV] évitez d'en reprendre une si vous tenez à la vie. ~ ❗❗ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.removeFromInv(" Potion de Poison")
			c.accesInventory()
		}
	}
	if foundPotion2 == false {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~  Vous n'avez pas de Potion de Poison, si besoin, allez chez le Marchand en acheter. ~ ❌ ]\n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		c.accesInventory()
	}
}

func (c *Character) end() {
	clearScreen()
	fmt.Print("\n\n\n[ 💬  ~ Voulez vous vraiment quitter le jeu ? ~ ] \n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Tapez [1] pour : Oui\n")
	fmt.Print("Tapez [2] pour : Non\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			clearScreen()
			os.Exit(3)
		} else if scanner.Text() == "2" {
			c.menu()
		} else {
			fmt.Println("Il faut que tu écrive 1 ou 2.")
		}
	}
}

func (c *Character) takePot() {
	clearScreen()
	fmt.Print("\n\n[ 💬  ~ Voulez-vous utiliser une Potion de Soin ? ~ ] ")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Tapez [1] pour : Oui\n")
	fmt.Print("Tapez [2] pour : Non\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.takePot2()
		} else if scanner.Text() == "2" {
			c.menu()
		} else {
			fmt.Println("Il faut que tu écrive 1 ou 2.")
		}
	}
}

func (c *Character) takePot2() {
	foundPotion := false
	for _, letters := range c.inv {
		if letters == " Potion de Soin" {
			foundPotion = true
			c.pv += 50
			if c.pv > c.pvmax {
				c.pv = c.pvmax
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous avez gagné [50 PV] et vous êtes actuellement à [", c.pv, " PV] ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				c.removeFromInv(" Potion de Soin")
				time.Sleep(3 * time.Second)
				c.accesInventory()
			} else {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous avez gagné [50 PV] et vous êtes actuellement à [", c.pv, " PV] ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.accesInventory()
			}
		}
	}
	if foundPotion == false {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~ Vous n'avez pas de Potion de Soin, si besoin, allez chez le Marchand en acheter. ~ ❌ ]\n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		c.accesInventory()
	}
}

func (c *Character) removeFromInv(item string) {
	var index int = -1
	for i := 0; i < len(c.inv) && index == -1; i++ {
		if c.inv[i] == item {
			index = i
		}
	}
	temp := []string{}
	temp = append(temp, c.inv[:index]...)
	temp = append(temp, c.inv[index+1:]...)
	c.inv = temp
}

func (c *Character) addFromInv(item string) {
	c.inv = append(c.inv, item)
}

func (c Character) merchant() {
	clearScreen()
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("💰 ~ Nombre de Pierres préciseuses actuelles : ")
	fmt.Print(c.gold, "💎")
	fmt.Print("\n\n")
	fmt.Print("Tapez [1] pour acheter : Potion de Poison (8💎)\n")
	fmt.Print("Tapez [2] pour acheter : Potion de Soin (5💎)\n")
	fmt.Print("Tapez [3] pour acheter : Cuir de Sanglier (12💎)\n")
	fmt.Print("Tapez [4] pour acheter : Batôn Ensorcellé (14💎)\n")
	fmt.Print("Tapez [5] pour acheter : Livre de Sort Magique (16💎)\n")
	fmt.Print("Tapez [6] pour acheter : Fourrure de Lion de Nuit (11💎)\n")
	fmt.Print("Tapez [7] pour acheter : Petit Lingot de Fer (8💎)\n\n")
	fmt.Print("(Tapez [8] pour retourner au Menu)\n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			if c.gold-8 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[❌ ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Potion de Poison")
				c.gold -= 8
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Potion de poison] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "2" {
			if c.gold-5 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Potion de Soin")
				c.gold -= 5
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Potion de soin] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "3" {
			if c.gold-12 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Cuir de Sanglier")
				c.gold -= 12
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Cuir de Sanglier] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "4" {
			if c.gold-14 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Batôn Ensorcellé")
				c.gold -= 14
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Batôn Ensorcellé] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "5" {
			if c.gold-16 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Livre de Sort Magique")
				c.gold -= 16
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Livre de Sort Magique] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "6" {
			if c.gold-11 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Fourrure de Lion de Nuit")
				c.gold -= 11
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Fourrure de Lion de Nuit] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "7" {
			if c.gold-8 < 0 {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses pour acheter cet objet. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			} else {
				c.addFromInv(" Petit Lingot de Fer")
				c.gold -= 8
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous venez d'acheter l'objet [Petit Lingot de Fer] et celui-ci a était ajouté à votre inventaire ! ~ ✔️ ]\n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.merchant()
			}
		} else if scanner.Text() == "8" {
			c.menu()
		} else {
			fmt.Print("Il faut que tu écrive 1, 2, 3, 4, 5, 6, 7 ou 8.")
		}
	}
}

func (c *Character) dead() {
	if c.pv <= 0 {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ 💀", " ~  Vous êtes mort ! ~ 💀 ]\n\n")
		fmt.Print("Tapez [1] pour Ressusciter\n")
		fmt.Print("Tapez [2] pour Quitter le jeu\n")
		fmt.Print("\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "1" {
				c.pv = 50
				c.relive()
			} else if scanner.Text() == "2" {
				os.Exit(3)
			} else {
				fmt.Println("Il faut que tu écrive 1 ou 2.")
			}
		}
	}
}

func (c *Character) relive() {
	clearScreen()
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("[ ❤️ ", " ~ Vous avez ressusciter, essayez de survivre la prochaine fois ! ~ ❤️ ] \n\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	time.Sleep(3 * time.Second)
	c.menu()
}

func isinslice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (c *Character) spellSkill() {
	foundSkill := false
	for _, letters := range c.skill {
		if letters == "Boule de Feu" {
			foundSkill = true
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ❌", " ~ Vous possédez déjà ce sort. ~ ❌ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.accesInventory()
		}
	}
	if foundSkill == false {
		clearScreen()
		c.removeFromInv(" Livre de Sort Magique")
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ✔️ ", " ~ Félicitation, vous venez d'acquérir le skill Boule de Feu ! ~ ✔️ ] \n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		if !(isinslice(c.skill, "Boule de Feu")) {
			c.skill = append(c.skill, "Boule de Feu")
		}
		c.accesInventory()
	}
}

func (c *Character) spellBook() {
	foundBook := false
	for _, letters := range c.inv {
		if letters == " Livre de Sort Magique" {
			foundBook = true
			c.spellSkill()
		}
	}

	if foundBook == false {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~ Vous n'avez pas de Livre de Sort Magique, si besoin, allez chez le Marchand en acheter. ~ ❌ ] \n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		c.accesInventory()
	}
}

func (c *Character) blacksmith() {
	clearScreen()
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("🔨 ~ Forgeron du village (Coût : 5💎 La Fabrication) :\n\n")
	fmt.Print("Tapez [1] pour Fabriquer un Batôn de Dégénérecensce (Batôn Ensorcellé + Potion de Poison) \n")
	fmt.Print("Tapez [2] pour Fabriquer un Manteau de Nuit (Fourrure de Lion de Nuit + Cuir de Sanglier)\n")
	fmt.Print("Tapez [3] pour Fabriquer un Casque de Fer (Petit Lingot de Fer + Cuir de Sanglier)\n")
	fmt.Print("\n(Tapez [4] pour retourner au Menu)")
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.craft1()
		} else if scanner.Text() == "2" {
			c.craft2()
		} else if scanner.Text() == "3" {
			c.craft3()
		} else if scanner.Text() == "4" {
			c.menu()
		} else {
			fmt.Println("Il faut que tu écrive 1, 2, 3 ou 4.")
		}
	}
}

func (c *Character) craft1() {
	if research(c.inv, " Batôn Ensorcellé") && research(c.inv, " Potion de Poison") {
		if c.gold < 5 {
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses. ~ ❌ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.blacksmith()
		} else {
			c.removeFromInv(" Batôn Ensorcellé")
			c.removeFromInv(" Potion de Poison")
			c.addFromInv(" Batôn de Dégénérecensce")
			c.gold -= 5
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ✔️ ", " ~ Félicitation, vous venez de fabriquer un Batôn de Dégénérecensce ! ~ ✔️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.blacksmith()
		}
	}

	if !(research(c.inv, " Batôn Ensorcellé") && research(c.inv, " Potion de Poison")) {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~ Vous n'avez pas les matériaux nécessaire, si besoin, allez chez le Marchand du village les acheter. ~ ❌ ] \n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		c.blacksmith()
	}
}

func (c *Character) craft2() {
	if research(c.inv, " Fourrure de Lion de Nuit") && research(c.inv, " Cuir de Sanglier") {
		if c.gold < 5 {
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses. ~ ❌ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.blacksmith()
		} else {
			c.removeFromInv(" Fourrure de Lion de Nuit")
			c.removeFromInv(" Cuir de Sanglier")
			c.addFromInv(" Manteau de Nuit")
			c.gold -= 5
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ✔️ ", " ~ Félicitation, vous venez de fabriquer un Manteau de Nuit ! ~ ✔️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.blacksmith()
		}
	}

	if !(research(c.inv, " Fourrure de Lion de Nuit") && research(c.inv, " Cuir de Sanglier")) {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~ Vous n'avez pas les matériaux nécessaire, si besoin, allez chez le Marchand du village les acheter. ~ ❌ ] \n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		c.blacksmith()
	}
}

func (c *Character) craft3() {
	if research(c.inv, " Petit Lingot de Fer") && research(c.inv, " Cuir de Sanglier") {
		if c.gold < 5 {
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ❌", " ~ Vous n'avez pas assez de Pierres précieuses. ~ ❌ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.blacksmith()
		} else {
			c.removeFromInv(" Petit Lingot de Fer")
			c.removeFromInv(" Cuir de Sanglier")
			c.addFromInv(" Casque de Fer")
			c.gold -= 5
			clearScreen()
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ✔️ ", " ~ Félicitation, vous venez de fabriquer un Casque de Fer ! ~ ✔️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			c.blacksmith()
		}
	}

	if !(research(c.inv, " Petit Lingot de Fer") && research(c.inv, " Cuir de Sanglier")) {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~ Vous n'avez pas les matériaux nécessaire, si besoin, allez chez le Marchand du village les acheter. ~ ❌ ] \n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
		c.blacksmith()
	}
}

func research(tab []string, s string) bool {
	for _, a := range tab {
		if a == s {
			return true
		}
	}
	return false
}

func (c *Character) armory() {
	clearScreen()
	fmt.Print("\n\n🗡️  ~ Armurerie :\n")
	fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	fmt.Print("Arme Principale : ", c.first, "\n")
	fmt.Print("Arme Secondaire : ", c.equipment.second, "\n")
	fmt.Print("Tête : ", c.equipment.head, "\n")
	fmt.Print("Torse : ", c.equipment.chest, "\n")
	fmt.Print("Pieds : ", "Bottines en Cuir \n")
	fmt.Print("\n(Tapez [1] pour vous équiper de votre Arme Secondaire)")
	fmt.Print("\n(Tapez [2] pour vous équiper de votre Equipement de Tête)")
	fmt.Print("\n(Tapez [3] pour vous équiper de votre Equipement de Torse)")
	fmt.Print("\n(Tapez [4] pour retourner au Menu)")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			if research(c.inv, " Batôn de Dégénérecensce") {
				c.equip(" Batôn de Dégénérecensce")
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Batôn de Dégénérecensce équipé ! ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.armory()
			} else {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez de Batôn de Dégénérecensce. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.armory()
			}
		} else if scanner.Text() == "2" {
			if research(c.inv, " Casque de Fer") {
				c.removeFromInv(" Casque de Fer")
				c.equipment.head = "Casque de Fer (+10 PV)"
				c.pvmax += 10
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Casque de Fer équipé ! ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.armory()
			} else {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez de Casque de Fer. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.armory()
			}
		} else if scanner.Text() == "3" {
			if research(c.inv, " Manteau de Nuit") {
				c.removeFromInv(" Manteau de Nuit")
				c.equipment.chest = "Manteau de Nuit (+25 PV)"
				c.pvmax += 25
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Manteau de Nuit équipé ! ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.armory()
			} else {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ❌", " ~ Vous n'avez de Manteau de Nuit. ~ ❌ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
				c.armory()
			}
		} else if scanner.Text() == "4" {
			c.menu()
		} else {
			fmt.Println("Il faut que tu écrive 1, 2, 3 ou 4.")
		}
	}
}

func (c *Character) equip(itemName string) {
	switch itemName {
	case " Batôn de Dégénérecensce":
		c.removeFromInv(" Batôn de Dégénérecensce")
		c.equipment.second = "Batôn de Dégénérecensce"
	}
}

func (c *Character) combat() {
	clearScreen()
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("[ ☠️ ", " ~ Êtes-vous sur de vouloir combattre le Boss ? ~ ☠️ ] \n\n")
	fmt.Print("\n(Tapez [1] pour Oui.)")
	fmt.Print("\n(Tapez [2] pour Non.)")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.boss()
			break
		} else if scanner.Text() == "2" {
			c.menu()
		} else {
			fmt.Println("Il faut que tu écrive 1 ou 2.")
		}
	}
}

func (c *Character) boss() {
	clearScreen()

	asciiArt :=

		`▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂





                 ___====-_  _-====___
           _--^^^#####//      \\#####^^^--_         
        _-^##########// (    ) \\##########^-_
       -############//  |\^^/|  \\############-
     _/############//   (@::@)   \\############\_
    /#############((     \\//     ))#############\
   -###############\\    (oo)    //###############-
  -#################\\  / VV \  //#################-
 -###################\\/      \//###################-
_#/|##########/\######(   /\   )######/\##########|\#_
|/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
'  |/  V  V  '   V  \#\| |  | |/#/  V   '  V  V  \|  '
   '   '  '      '   / | |  | | \   '      '  '   '
                    (  | |  | |  )
                   __\ | |  | | /__
                  (vvv(VVV)(VVV)vvv)
			 

					
▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂


[ 💬 ~ GARAKROND le Dragon Noir apparaît !!! ~ ]
 `
	fmt.Print(asciiArt)
	time.Sleep(3 * time.Second)
	c.bossbar()
}

func (c *Character) bossbar() {
	clearScreen()
	sBar := "█"
	sVide := "                                        "
	for i := 0; i <= 75; i++ {
		time.Sleep(20 * time.Millisecond)
		clearScreen()
		if i%2 == 0 {
			sBar += "█"
			sVide = sVide[:len(sVide)-1]

		}

		fmt.Print(`▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂


	[` + sBar + sVide + `PV de GARAKROND ]`)
		fmt.Print(`


                 ___====-_  _-====___
           _--^^^#####//      \\#####^^^--_         
        _-^##########// (    ) \\##########^-_
       -############//  |\^^/|  \\############-
     _/############//   (@::@)   \\############\_
    /#############((     \\//     ))#############\     
   -###############\\    (oo)    //###############-
  -#################\\  / VV \  //#################-
 -###################\\/      \//###################-
_#/|##########/\######(   /\   )######/\##########|\#_
|/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
'  |/  V  V  '   V  \#\| |  | |/#/  V   '  V  V  \|  '
   '   '  '      '   / | |  | | \   '      '  '   '
                    (  | |  | |  )
                   __\ | |  | | /__
                  (vvv(VVV)(VVV)vvv)
    
           
				  
▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂


[ 💬  ~ Bonne chance pour le tuer !!! ~ ]`)
	}
	time.Sleep(3 * time.Second)
	var Dragon Monstre
	Dragon.initDragon("GARAKRON", 160, 7)
	Dragon.trainingFight(c)
}

func (c *Character) invfight(m *Monstre) {
	clearScreen()
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("[ 💼 ", " ~ Inventaire ~ 💼 ] \n\n")
	fmt.Print("Objets : ", c.inv)
	fmt.Print("\n\n(Tapez [1] pour utiliser une Potion de Soin.)")
	fmt.Print("\n(Tapez [2] pour utiliser votre Batôn de Dégénérecensce.)")
	fmt.Print("\n(Tapez [3] pour retourner au Menu du Combat.)")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.takePot3()
			break
		} else if scanner.Text() == "2" {
			m.batondegene(c)
			break
		} else if scanner.Text() == "3" {
			break
		} else {
			fmt.Println("Il faut que tu écrive 1, 2 ou 3.")
		}
	}
}

func (m *Monstre) initDragon(nom string, pvmax int, attaque int) {
	m.nom = nom
	m.pv = pvmax / 2
	m.pvmax = pvmax
	m.attaque = attaque
}
func (m *Monstre) dragonPattern(c *Character, turn int) {
	damage := m.attaque
	if turn != 0 && turn%3 == 0 {
		damage = m.attaque * 2
	}
	c.pv -= damage
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("[ 💥️ ~ " + m.nom + " t'attaque de " + strconv.Itoa(damage) + " Dégats sur toi, il te reste " + strconv.Itoa(c.pv) + "/" + strconv.Itoa(c.pvmax) + " ~ 💥️ ] \n")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	time.Sleep(3 * time.Second)
}

func (m *Monstre) trainingFight(c *Character) {
	for i := 0; i < 1000; i++ {
		clearScreen()
		barrePV(m)
		partDragon()
		c.attaque1(m)
		if m.pv < 1 {
			c.win()
			break
		}
		barrePV(m)
		partDragon()
		m.dragonPattern(c, i)
		c.dead()
	}
}

func (c *Character) attaque1(m *Monstre) {

	fmt.Print("\n PV de ", c.nom, " : ", c.pv)
	fmt.Print("\n\n[ 💬  ~ Menu du Combat ~ ]\n")
	fmt.Print("\n(Tapez [1] pour Attaquer.)")
	fmt.Print("\n(Tapez [2] pour Accéder à l'Inventaire.)")
	fmt.Print("\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			m.pv -= 10
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ 💥️ ", " ~ Tu attaque de 10 de Dégats sur le ", m.nom, " celui-ci passe à ", m.pv, " ~ 💥️ ] \n")
			fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(3 * time.Second)
			break
		} else if scanner.Text() == "2" {
			c.invfight(m)
			break
		} else {
			fmt.Println("Il faut que tu écrive 1 ou 2.")
		}
	}
	clearScreen()
}

func partDragon() {
	fmt.Print(`

	 
                 ___====-_  _-====___
           _--^^^#####//      \\#####^^^--_
        _-^##########// (    ) \\##########^-_
       -############//  |\^^/|  \\############-
     _/############//   (@::@)   \\############\_
    /#############((     \\//     ))#############\       
   -###############\\    (oo)    //###############-
  -#################\\  / VV \  //#################-
 -###################\\/      \//###################-
_#/|##########/\######(   /\   )######/\##########|\#_
|/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
'  |/  V  V  '   V  \#\| |  | |/#/  V   '  V  V  \|  '
   '   '  '      '   / | |  | | \   '      '  '   '
                    (  | |  | |  )
                   __\ | |  | | /__
                  (vvv(VVV)(VVV)vvv)



▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂

`)

}

func barrePV(m *Monstre) {
	fmt.Print(`▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂
	
	
	`)
	if m.pv <= 0 {
		fmt.Print("[                                         PV de GARAKROND ]")
	}
	if m.pv > 0 && m.pv <= 2 {
		fmt.Print("[█                                        PV de GARAKROND ]")
	}
	if m.pv > 2 && m.pv <= 4 {
		fmt.Print("[██                                       PV de GARAKROND ]")
	}
	if m.pv > 4 && m.pv <= 6 {
		fmt.Print("[███                                      PV de GARAKROND ]")
	}
	if m.pv > 6 && m.pv <= 8 {
		fmt.Print("[████                                     PV de GARAKROND ]")
	}
	if m.pv > 8 && m.pv <= 10 {
		fmt.Print("[█████                                    PV de GARAKROND ]")
	}
	if m.pv > 10 && m.pv <= 12 {
		fmt.Print("[██████                                   PV de GARAKROND ]")
	}
	if m.pv > 12 && m.pv <= 14 {
		fmt.Print("[███████                                  PV de GARAKROND ]")
	}
	if m.pv > 14 && m.pv <= 16 {
		fmt.Print("[████████                                 PV de GARAKROND ]")
	}
	if m.pv > 16 && m.pv <= 18 {
		fmt.Print("[█████████                                PV de GARAKROND ]")
	}
	if m.pv > 18 && m.pv <= 20 {
		fmt.Print("[██████████                               PV de GARAKROND ]")
	}
	if m.pv > 20 && m.pv <= 22 {
		fmt.Print("[███████████                              PV de GARAKROND ]")
	}
	if m.pv > 22 && m.pv <= 24 {
		fmt.Print("[████████████                             PV de GARAKROND ]")
	}
	if m.pv > 24 && m.pv <= 26 {
		fmt.Print("[█████████████                            PV de GARAKROND ]")
	}
	if m.pv > 26 && m.pv <= 28 {
		fmt.Print("[██████████████                           PV de GARAKROND ]")
	}
	if m.pv > 28 && m.pv <= 30 {
		fmt.Print("[███████████████                          PV de GARAKROND ]")
	}
	if m.pv > 30 && m.pv <= 32 {
		fmt.Print("[████████████████                         PV de GARAKROND ]")
	}
	if m.pv > 32 && m.pv <= 34 {
		fmt.Print("[█████████████████                        PV de GARAKROND ]")
	}
	if m.pv > 34 && m.pv <= 36 {
		fmt.Print("[██████████████████                       PV de GARAKROND ]")
	}
	if m.pv > 36 && m.pv <= 38 {
		fmt.Print("[███████████████████                      PV de GARAKROND ]")
	}
	if m.pv > 38 && m.pv <= 40 {
		fmt.Print("[████████████████████                     PV de GARAKROND ]")
	}
	if m.pv > 40 && m.pv <= 42 {
		fmt.Print("[█████████████████████                    PV de GARAKROND ]")
	}
	if m.pv > 42 && m.pv <= 44 {
		fmt.Print("[██████████████████████                   PV de GARAKROND ]")
	}
	if m.pv > 44 && m.pv <= 46 {
		fmt.Print("[███████████████████████                  PV de GARAKROND ]")
	}
	if m.pv > 46 && m.pv <= 48 {
		fmt.Print("[████████████████████████                 PV de GARAKROND ]")
	}
	if m.pv > 48 && m.pv <= 50 {
		fmt.Print("[█████████████████████████                PV de GARAKROND ]")
	}
	if m.pv > 50 && m.pv <= 52 {
		fmt.Print("[██████████████████████████               PV de GARAKROND ]")
	}
	if m.pv > 52 && m.pv <= 54 {
		fmt.Print("[███████████████████████████              PV de GARAKROND ]")
	}
	if m.pv > 54 && m.pv <= 56 {
		fmt.Print("[████████████████████████████             PV de GARAKROND ]")
	}
	if m.pv > 56 && m.pv <= 58 {
		fmt.Print("[█████████████████████████████            PV de GARAKROND ]")
	}
	if m.pv > 58 && m.pv <= 60 {
		fmt.Print("[██████████████████████████████           PV de GARAKROND ]")
	}
	if m.pv > 60 && m.pv <= 62 {
		fmt.Print("[███████████████████████████████          PV de GARAKROND ]")
	}
	if m.pv > 62 && m.pv <= 64 {
		fmt.Print("[████████████████████████████████         PV de GARAKROND ]")
	}
	if m.pv > 64 && m.pv <= 66 {
		fmt.Print("[█████████████████████████████████        PV de GARAKROND ]")
	}
	if m.pv > 66 && m.pv <= 68 {
		fmt.Print("[██████████████████████████████████       PV de GARAKROND ]")
	}
	if m.pv > 68 && m.pv <= 70 {
		fmt.Print("[███████████████████████████████████      PV de GARAKROND ]")
	}
	if m.pv > 70 && m.pv <= 72 {
		fmt.Print("[████████████████████████████████████     PV de GARAKROND ]")
	}
	if m.pv > 72 && m.pv <= 74 {
		fmt.Print("[█████████████████████████████████████    PV de GARAKROND ]")
	}
	if m.pv > 74 && m.pv <= 76 {
		fmt.Print("[██████████████████████████████████████   PV de GARAKROND ]")
	}
	if m.pv > 76 && m.pv <= 78 {
		fmt.Print("[███████████████████████████████████████  PV de GARAKROND ]")
	}
	if m.pv > 78 && m.pv <= 80 {
		fmt.Print("[███████████████████████████████████████  PV de GARAKROND ]")
	}
}

func (m *Monstre) batondegene(c *Character) {
	foundBaton := false
	for _, letters := range c.inv {
		if letters == " Batôn de Dégénérecensce" {
			foundBaton = true
			clearScreen()
			m.pv -= 5
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ☠️ ", "  ~ Aie...Le Batôn lance un sort de Poison sur GARAKRON, celui-ci perds 10 de PV et passe à [", m.pv, " PV] ! ~ ☠️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(1 * time.Second)
			m.pv -= 5
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ☠️ ", "  ~ Aie...Le Batôn lance un sort de Poison sur GARAKRON, celui-ci perds 10 de PV et passe à [", m.pv, " PV] ! ~ ☠️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(1 * time.Second)
			m.pv -= 5
			fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			fmt.Print("[ ☠️ ", "  ~ Aie...Le Batôn lance un sort de Poison sur GARAKRON, celui-ci perds 10 de PV et passe à [", m.pv, " PV] ! ~ ☠️ ] \n\n\n")
			fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
			time.Sleep(1 * time.Second)

		}
	}
	if foundBaton == false {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~  Vous n'avez pas de Batôn de Dégénérecensce. ~ ❌ ]\n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
	}
}

func (c *Character) takePot3() {
	foundPotion := false
	for _, letters := range c.inv {
		if letters == " Potion de Soin" {
			foundPotion = true
			c.pv += 50
			if c.pv > c.pvmax {
				c.pv = c.pvmax
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous avez gagné [50 PV] et vous êtes actuellement à [", c.pv, " PV] ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				c.removeFromInv(" Potion de Soin")
				time.Sleep(3 * time.Second)
			} else {
				clearScreen()
				fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				fmt.Print("[ ✔️ ", " ~ Vous avez gagné [50 PV] et vous êtes actuellement à [", c.pv, " PV] ~ ✔️ ] \n\n\n")
				fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
				time.Sleep(3 * time.Second)
			}
		}
	}
	if foundPotion == false {
		clearScreen()
		fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		fmt.Print("[ ❌", " ~ Vous n'avez pas de Potion de Soin, si besoin, allez chez le Marchand en acheter. ~ ❌ ]\n\n\n")
		fmt.Print("▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
		time.Sleep(3 * time.Second)
	}
}

func (c *Character) win() {

	clearScreen()
	fmt.Print("\n\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	fmt.Print("[ 🎉 ", " ~ Félicitation, vous avez vaincu GARAKRON, le Dragon !!! ~ 🎉 ] \n\n")
	fmt.Print("\n(Tapez [1] pour Retourner au Menu.)")
	fmt.Print("\n(Tapez [2] pour Quitter le Jeu)")
	fmt.Print("\n\n▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂▂\n\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "1" {
			c.menu()
		} else if scanner.Text() == "2" {
			c.end()
		} else {
			fmt.Println("Il faut que tu écrive 1, 2 ou 3.")
		}
	}
}
