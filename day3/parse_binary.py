def create_two_dimensional_array(file):
    two_dimensional_array = []
    with open(file) as fe:
        for line in fe:
            two_dimensional_array.append([str(i) for i in line.strip()])
    fe.close()
    return two_dimensional_array


def calculate_gamma_rate(two_dimensional_binary_array):
    gamma = ""
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
        else:
            gamma += "0"
        zeroes = 0
        ones = 0
    return gamma


def reverse_binary(binary_string: str):
    reversed_binary_string = ""
    for i in binary_string:
        if i == "0":
            reversed_binary_string += "1"
        if i == "1":
            reversed_binary_string += "0"
    return reversed_binary_string


def main():
    two_dimensional_binary_array = create_two_dimensional_array("input.txt")
    gamma = calculate_gamma_rate(two_dimensional_binary_array)
    epsilon = reverse_binary(gamma)

    print(f'Power consumption: {int(epsilon, 2) * int(gamma, 2)}')


main()
