# Input is two lists, goal is to find the total distance between the two
# Total distance = difference in pairs, smallest to largest
# E.g. smallest from first list and smallest from second list, etc.

# Easiest way would be to read in and sort both lists, then go through each
# and sum up the differences

left_list: list[int] = []
right_list: list[int] = []

def read_input_file(fileName: str) -> None:
    f = open(fileName, 'r') 
    for l in f:
        line = l.strip().split()
        left_list.append(int(line[0]))
        right_list.append(int(line[1]))

    f.close()

def list_distance() -> None:
    read_input_file('input.txt')
    left_list.sort()
    right_list.sort()

    sum = 0
    for i in range(len(left_list)):
        sum += abs(left_list[i] - right_list[i])

    print(sum)

# Second half wants a similarity score. In other words, how often numbers in 
# the left list appear in the right list.
# We can use a dictionary to keep track of number of occurrances in the right list,
# and then iterate through the left list and sum up number of occurrances in right list

def similarity_score() -> None:
    read_input_file('input.txt')
    d: dict = {}
    sim_score = 0

    # Need to find sum on numbers in right list, then read from left list
    for r in right_list:
        if r in d:
            d[r] += 1
        else:
            d[r] = 1

    for l in left_list:
        if l in d:
            sim_score += l * d[l]

    print(sim_score)


if __name__ == '__main__':
    # list_distance()
    similarity_score()
