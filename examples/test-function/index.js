exports.handler = async function(args = []) {
    const input = args[0] || "World";
    
    // Simulate some processing
    await new Promise(resolve => setTimeout(resolve, 100));
    
    return {
        message: `Hello, ${input}!`,
        timestamp: new Date().toISOString(),
        processedAt: Date.now(),
        inputReceived: args
    };
};