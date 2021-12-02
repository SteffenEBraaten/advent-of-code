def caluclateDepthDeprecated(file):
    horizontalPos = 0
    verticalPos = 0
    with open(file, 'r') as fe:
        for command in open(file):
            commandList = command.split()
            if commandList[0] == "forward":
                horizontalPos += int(commandList[1])
            if commandList[0] == "down":
                verticalPos += int(commandList[1])
            if commandList[0] == "up":
                verticalPos -= int(commandList[1])
    fe.close()
    return horizontalPos * verticalPos

def calculateDept(file):
    aim = 0
    verticalPos = 0
    horizontalPos = 0
    with open(file, 'r') as fe:
        for command in open(file):
            commandList = command.split()
            if commandList[0] == "forward":
                horizontalPos += int(commandList[1])
                if aim != 0:
                    verticalPos += (aim * int(commandList[1]))
            if commandList[0] == "down":
                aim += int(commandList[1])
            if commandList[0] == "up":
                aim -= int(commandList[1])
    fe.close()
    return horizontalPos * verticalPos

def main():
    file = 'input.txt'
    deprecatedDepth = caluclateDepthDeprecated(file)
    print(f'Deprecated depth calculation: {deprecatedDepth}\n')

    depth = calculateDept(file)
    print(f'Calculated depth: {depth}')

main()
