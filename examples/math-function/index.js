exports.handler = async function(args = []) {
    const operation = args[0] || "add";
    const numbers = args.slice(1).map(n => parseFloat(n)).filter(n => !isNaN(n));
    
    if (numbers.length === 0) {
        throw new Error("No valid numbers provided");
    }
    
    let result;
    switch (operation) {
        case "add":
            result = numbers.reduce((sum, num) => sum + num, 0);
            break;
        case "multiply":
            result = numbers.reduce((product, num) => product * num, 1);
            break;
        case "max":
            result = Math.max(...numbers);
            break;
        case "min":
            result = Math.min(...numbers);
            break;
        case "average":
            result = numbers.reduce((sum, num) => sum + num, 0) / numbers.length;
            break;
        default:
            throw new Error(`Unsupported operation: ${operation}`);
    }
    
    return {
        operation,
        numbers,
        result,
        timestamp: new Date().toISOString()
    };
};