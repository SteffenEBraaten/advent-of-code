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


def delete_multiple(list, indices):
    indices = sorted(indices, reverse=True)
    for idx in indices:
        if idx < len(list):
            list.pop(idx)


def calculate_oxygen(two_dimensional_binary_array):
    oxygen = [row[:] for row in two_dimensional_binary_array]
    for i in range(len(oxygen[0])):
        zeroes = []
        ones = []
        for j in range(len(oxygen)):
            if oxygen[j][i] == "0":
                zeroes.append(j)
            else:
                ones.append(j)
        if len(oxygen) == 1:
            return "".join(oxygen[0])
        if len(ones) >= len(zeroes):
            delete_multiple(oxygen, zeroes)
        else:
            delete_multiple(oxygen, ones)
    return "".join(oxygen[0])


def calculate_c02(two_dimensional_binary_array):
    c02 = [row[:] for row in two_dimensional_binary_array]
    for i in range(len(c02[0])):
        zeroes = []
        ones = []
        for j in range(len(c02)):
            print(j)
            if c02[j][i] == "0":
                zeroes.append(j)
            else:
                ones.append(j)
        if len(c02) == 1:
            return "".join(c02[0])
        if len(ones) >= len(zeroes):
            delete_multiple(c02, ones)
        else:
            delete_multiple(c02, zeroes)
    return "".join(c02[0])


def main():
    two_dimensional_binary_array = create_two_dimensional_array("input.txt")
    gamma, epsilon = calculate_gamma_and_epsilon(two_dimensional_binary_array)

    print(f'Power consumption: {int(epsilon, 2) * int(gamma, 2)}')

    oxygen = calculate_oxygen(two_dimensional_binary_array)
    c02 = calculate_c02(two_dimensional_binary_array)

    print(f'Life support rating: {int(oxygen, 2) * int(c02, 2)}')


main()
