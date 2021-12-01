# Day one of Advent of Code
import sys

def __read_file_of_numbers(file):
    try:
        with open(file, 'r') as fe:
            print('Reading from file...')
            number_list = fe.read().splitlines()
            fe.close()
            print('Done reading from file!')
            number_list = list(map(int, number_list)) # Turning list of strings into list of ints
            return number_list
    except IOError as error:
        print(error)
        sys.exit()
        

def count_larger_than_previous_measurement(listOfNumbers):
    number_of_larger_than_previous = 0
    for index, number in enumerate(listOfNumbers):
        if index != 0 and number > listOfNumbers[index - 1]:
            number_of_larger_than_previous = number_of_larger_than_previous + 1
    return number_of_larger_than_previous

def create_three_measurement_sliding_window(listOfNumbers):
    numberList = []
    for index, number in enumerate(listOfNumbers):
        if (index + 2) < len(listOfNumbers):
            numberList.append(number + listOfNumbers[index + 1] + listOfNumbers[index + 2])
    return numberList


def main():
    file = 'input.txt'
    listOfNumbers = __read_file_of_numbers(file)
    sum = count_larger_than_previous_measurement(listOfNumbers)
    print(f'There are {sum} measueres that are larger than its previous measure')
    
    three_measurement_sliding_window_list = create_three_measurement_sliding_window(listOfNumbers)
    three_measurement_sliding_sum = count_larger_than_previous_measurement(three_measurement_sliding_window_list)
    print(f'There are {three_measurement_sliding_sum} three-measueres that are larger than its previous measure')

main()