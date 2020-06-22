package humansrv

import (
	"errors"
	"fmt"
	"strings"

	"github.com/matiasresca/xmen-meli/pkg/core/domain"
	"github.com/matiasresca/xmen-meli/pkg/core/ports"
)

var (
	lettersAvaible           = []byte{'A', 'T', 'C', 'G'}
	checkTypeHorizonal       = "horizontal"
	checkTypeVertical        = "vertical"
	checkTypeDiagonal        = "diagonal"
	checkTypeReverseDiagonal = "reverse_diagonal"
)

type service struct {
	repo ports.HumanRepository
}

func NewService(repo ports.HumanRepository) ports.HumanService {
	return &service{
		repo: repo,
	}
}

func (s *service) GetStats() (*domain.Stats, error) {
	stats := domain.Stats{}
	//Busco si fue procesado anteriormente.-
	humans, err := s.repo.GetAll()
	if err != nil {
		fmt.Println(err)
		return &stats, err
	}
	for _, human := range humans {
		if human.IsMutant {
			stats.CountMutantDna++
		} else {
			stats.CountHumanDna++
		}
		if stats.CountHumanDna > 0 {
			stats.Ratio = float64(stats.CountMutantDna) / float64(stats.CountHumanDna)
		}
	}
	return &stats, nil
}

func (s *service) CheckMutant(dna []string) (bool, error) {
	//Valido que no sea vacio.-
	if len(dna) == 0 {
		return false, errors.New("El DNA no puede ser vacio")
	}
	//Tamaño de filas el ADN.-
	dnaLen := len(dna)
	for _, letters := range dna {
		//Valido que las columnas sean iguales que las filas, la matriz sea NxN.-
		if dnaLen != len(letters) {
			return false, errors.New("El ADN no cumple con una matriz NxN")
		}
		//Valido que las letras sean compatibles.-
		for _, letter := range letters {
			validate := false
			for _, letterAvaible := range lettersAvaible {
				if byte(letter) == letterAvaible {
					validate = true
				}
			}
			if !validate {
				return false, errors.New("Las letras del ADN solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN")
			}
		}
	}
	//Busco si fue procesado anteriormente.-
	human, err := s.repo.GetByDna(dna)
	if err != nil {
		fmt.Println(err)
	}
	//Si ya fue procesado, devuelvo el valor seteado en el registro.-
	if human != nil {
		return human.IsMutant, nil
	}

	//Verifico si es mutante.-
	isMutant := s.isMutant(dna)

	//Guardo el DNA procesado.-
	err = s.repo.Save(domain.Human{Dna: dna, IsMutant: isMutant})
	if err != nil {
		panic(err)
	}
	return isMutant, nil
}

func (s *service) isMutant(dna []string) bool {
	//Algoritmo para validar si es mutante.-
	//Cantidad minima de caracteres consecutivos.-
	countDna := 4
	//Armo matriz de ADN.-
	sizeTable := len(dna)
	var matriz [][]string
	matriz = make([][]string, sizeTable)
	for k, v := range dna {
		dnaSplit := strings.Split(v, "")
		for _, dnaSplitValue := range dnaSplit {
			matriz[k] = append(matriz[k], dnaSplitValue)
		}
	}
	//Inicializo resultado.-
	result := false
	//Por cada valor de la matriz verifico consecutivos de manera horizonal, vertical, digonal y diagonal inversa a travez de una funcion recursiva.-
	for row := 0; row < sizeTable; row++ {
		for column := 0; column < sizeTable; column++ {
			if s.checkConsecutive(matriz, column, row, checkTypeHorizonal) >= countDna ||
				s.checkConsecutive(matriz, column, row, checkTypeVertical) >= countDna ||
				s.checkConsecutive(matriz, column, row, checkTypeDiagonal) >= countDna ||
				s.checkConsecutive(matriz, column, row, checkTypeReverseDiagonal) >= countDna {
				result = true
				break
			}
		}
		if result {
			break
		}
	}
	return result
}

func (s *service) checkConsecutive(matriz [][]string, column int, row int, checkType string) int {
	result := 1
	valueFind := matriz[row][column]
	//Segun el tipo de verificacion se aplica distinta estrategia de comprobacion.-
	exp := false
	switch checkType {
	case checkTypeHorizonal:
		column += 1
		exp = column < len(matriz[row])
	case checkTypeVertical:
		row += 1
		exp = row < len(matriz)
	case checkTypeDiagonal:
		row += 1
		column += 1
		exp = row < len(matriz) && column < len(matriz[row])
	case checkTypeReverseDiagonal:
		row += 1
		column -= 1
		exp = row < len(matriz) && column >= 0
	}
	//Valido la expresion y que sea iguales los valores.-
	if exp && valueFind == matriz[row][column] {
		result += s.checkConsecutive(matriz, column, row, checkType)
	}
	return result
}
