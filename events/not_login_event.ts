


class NotLoginEvent {
    ins: any[] = []

    constructor() {

    }

    on (callback: () => void) {
        this.ins.push(callback)
    }

    tigger () {
        for(var i in this.ins) {
            this.ins[i]()
        }
    }
    removeAll () {
        this.ins = []
    }
}


const notLoginEventIns = new NotLoginEvent()

export default {
    tigger: () => notLoginEventIns.tigger(),
    on: (callback: () => void) => notLoginEventIns.on(callback),
    removeAll: () => notLoginEventIns.removeAll()
}
