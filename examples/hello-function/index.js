exports.handler = async function(args = []) {
    const name = args[0] || "World";
    const count = args[1] || 1;
    
    const messages = [];
    for (let i = 0; i < count; i++) {
        messages.push(`Hello, ${name}! (${i + 1})`);
    }
    
    return {
        message: "Function executed successfully",
        results: messages,
        timestamp: new Date().toISOString(),
        inputArgs: args
    };
};