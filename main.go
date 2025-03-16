package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Mapping van naam naar level uit Input 2
	levelMap := make(map[string]int)

	// Stap 1: Lees Input 2 en vul de mapping
	file2, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Fout bij openen van input2.txt:", err)
		return
	}
	defer file2.Close()

	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		line := scanner2.Text()
		parts := strings.Split(line, "   ") // Drie spaties als scheiding
		if len(parts) < 2 {
			continue
		}
		var level int
		_, err := fmt.Sscanf(parts[1], "%d", &level)
		if err != nil {
			continue
		}
		name := strings.TrimSpace(parts[0])
		levelMap[name] = level
	}

	// Stap 2: Lees Input 1 en pas levels aan
	file1, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Fout bij openen van input1.txt:", err)
		return
	}
	defer file1.Close()

	scanner1 := bufio.NewScanner(file1)
	for scanner1.Scan() {
		line := scanner1.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 3 {
			fmt.Println(line) // Ongeldige regel, ongewijzigd uitvoeren
			continue
		}
		name := parts[0]
		// Controleer of de naam in de mapping zit
		if newLevel, exists := levelMap[name]; exists {
			// Level aanpassen
			fmt.Printf("%s:%d:%s\n", name, newLevel, parts[2])
		} else {
			// Geen match, originele regel behouden
			fmt.Println(line)
		}
	}

	if err := scanner1.Err(); err != nil {
		fmt.Println("Fout bij lezen van input1.txt:", err)
	}
	if err := scanner2.Err(); err != nil {
		fmt.Println("Fout bij lezen van input2.txt:", err)
	}
}
