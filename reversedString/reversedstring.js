function reversestring (inputstr) {
    let result = "";
    for (let character of inputstr){
        result = character + result;
    }
    return result;
}

console.log(reversestring('sdrt'));