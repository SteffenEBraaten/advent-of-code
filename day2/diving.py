def caluclateDepth(file):
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


def main():
    file = 'input.txt'
    depth = caluclateDepth(file)
    print(f'Depth: {depth}')

main()
