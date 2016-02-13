export function getTimestamp() {
    return new Date().toISOString();
}

export function log(...args: any[]) {
    if (!console || !console.log) {
        return;
    }

    args.unshift(getTimestamp());
    console.log.apply(console, args);
}
