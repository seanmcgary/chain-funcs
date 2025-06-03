
exports.handler = async function(args = []) {
	console.log('Hello world!');
	console.log('Arguments:', args);

	return { message: 'this is a message' };
}
