def create_two_dimensional_array(file):
    two_dimensional_array = []
    with open(file) as fe:
        for line in fe:
            two_dimensional_array.append([str(i) for i in line.strip()])
    return two_dimensional_array


def calculate_gamma_and_epsilon(two_dimensional_binary_array):
    gamma = ""
    epsilon = ""
    for i in range(len(two_dimensional_binary_array[0])):
        zeroes = 0
        ones = 0
        for j in range(len(two_dimensional_binary_array)):
            if two_dimensional_binary_array[j][i] == "0":
                zeroes += 1
            else:
                ones += 1
        if zeroes < ones:
            gamma += "1"
            epsilon += "0"
        else:
            gamma += "0"
            epsilon += "1"
        zeroes = 0
        ones = 0
    return gamma, epsilon


def main():
    two_dimensional_binary_array = create_two_dimensional_array("input.txt")
    gamma, epsilon = calculate_gamma_and_epsilon(two_dimensional_binary_array)

    print(f'Power consumption: {int(epsilon, 2) * int(gamma, 2)}')


main()
