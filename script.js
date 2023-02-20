const {readFile, writeFile} = require("fs/promises")

const main = async () => {
    const newIds = JSON.parse(await readFile("new_.json", 'utf8'));
    const oldData = JSON.parse(await readFile("exercise_muscle.json", 'utf8'))

    let result = [];

    const findNewId = oldMuscleId => {
        const data = newIds.find(({id}) => id === oldMuscleId)
        return {id:data.id, newId: data['new id']}
    }

    const findOldId = oldId => {
        return oldData.find( data => data.muscle_id === oldId )
    }

    const replaceId = (old, new_) => {
        if(Array.isArray(new_.newId)){
            const result = []
            new_.newId.forEach( id => {
                result.push({
                    ...old,
                    muscle_id:id
                })
            })
            // console.log({result})
            return result
        }

        if(old.id == 10138) {
            console.log("aqui",{
                ...old,
                muscle_id:new_.newId
            })
        }
        return {
            ...old,
            muscle_id:new_.newId
        }
    }

    oldData.forEach((old, i) => {
        
        const new_ = findNewId(old.muscle_id)
        
        const newData = replaceId(old, new_)
        if(i == 0) {
            console.log({old,new_,newData})
        }
        if(Array.isArray(newData)) {
            return result.push(...newData)
        } 
        result.push(newData)
    })

    result = result.flat().map((data,i) => {
        const {muscle_id, role, exercise_id} = data
        return {muscle_id, role: role.trim(), exercise_id} 
    })

    let resultClean = [];

    for (const data of result) {
        const {muscle_id, role, exercise_id} = data
        const hasALready = resultClean.find(data => {
            return muscle_id === data.muscle_id && role === data.role && exercise_id === data.exercise_id;
        })
        if(!hasALready) resultClean.push({muscle_id, role, exercise_id})
    }


    await writeFile("new.json", JSON.stringify(resultClean));

    result = resultClean.reduce((acc,data) => {
        const {muscle_id, role, exercise_id} = data
        acc+=`, (${muscle_id}, ${exercise_id}, '${role}')`
        return acc
    }, 'INSERT INTO exercise_musclePortion (muscle_id, exercise_id, role) VALUES ')

    await writeFile("new.sql", result)
    
}

main()