def main(params):
    examOneMarks = params.get("examOneMarks")
    examTwoMarks = params.get("examTwoMarks")
    
    fullMarks = examOneMarks + examTwoMarks
    averageMarks =  fullMarks / 2
        
    return {"averageMarks": averageMarks}
