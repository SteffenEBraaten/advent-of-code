# Day one of Advent of Code
import sys

def __read_file_of_numbers(file):
    try:
        with open(file, 'r') as fe:
            print('Reading from file...')
            number_list = fe.read().splitlines()
            fe.close()
            print('Done reading from file!')
            number_list = list(map(int, number_list)) # Truning it into list of ints
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

def main():
    file = 'input.txt'
    sum = count_larger_than_previous_measurement(__read_file_of_numbers(file))
    print(f'There are {sum} measueres that are larger than its previous measure')

main()