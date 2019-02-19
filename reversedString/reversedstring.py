def reversedstring(inputstr):

    result = ""
    for char in inputstr:
        result = char + result
    return result


print(reversedstring('dbfjbc'))
