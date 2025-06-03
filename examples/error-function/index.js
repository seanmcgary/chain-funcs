exports.handler = async function(args = []) {
    const shouldError = args[0] === "error";
    
    if (shouldError) {
        throw new Error("This is a test error");
    }
    
    return {
        success: true,
        message: "Function executed without error",
        args: args
    };
};