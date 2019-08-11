function main(params) {
    var examResult = '';
 
    if (params.examMarkes < 0 || params.examMarkes > 100) {
        examResult = 'ERROR: invalid exam mark';
    } else if (params.examMarkes >= 60) {
        examResult = 'Pass';
    } else {
	  examResult = 'Fail';
    }
	
    return { result: examResult };
}
