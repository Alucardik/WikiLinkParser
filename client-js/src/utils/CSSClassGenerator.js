export default class cn {
    constructor(blockName) {
        this._root = blockName;
    }

    root = () => {
        return this._root;
    }


    elem = (elemName, modName = undefined, val = undefined) => {
        let baseElem = `${this._root}__${elemName}`

        if (modName) {
            baseElem += `_${modName}`;

            if (val) {
                baseElem += `_${val}`;
            }
        }

        return baseElem;
    }

    mix = (mixElem) => {
        return new cn(`${mixElem} ${this._root}`)
    }
}
