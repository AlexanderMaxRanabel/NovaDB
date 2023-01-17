class NovoDB {
    private data: { [key: string]: any } = {};

    public insert(key: string, value: any) {
        this.data[key] = value;
    }

    public find(key: string) {
        return this.data[key];
    }

    public update(key: string, value: any) {
        this.data[key] = value;
    }

    public remove(key: string) {
        delete this.data[key];
    }
}

const db = new NovaDB();

db.insert("user1", { name: "John Doe", age: 30 });
console.log(db.find("user1")); // { name: "John Doe", age: 30 }

db.update("user1", { name: "Jane Doe", age: 35 });
console.log(db.find("user1")); // { name: "Jane Doe", age: 35 }

db.remove("user1");
console.log(db.find("user1")); // undefined
