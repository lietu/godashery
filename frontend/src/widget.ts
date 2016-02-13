export interface Widget {
    update(data: WidgetData): void;
}

export class BaseWidget {
    private id: string;

    constructor(id: string) {
        this.id = id;
    }
}

export interface WidgetData {
    type: string;
    id: string;
}

export interface WidgetList {
    [index: string]: Widget;
}

interface WidgetClassList {
    [index: string]: new (id: string) => Widget;
}

class WidgetManager {
    private widgets: WidgetClassList = {};

    public registerWidget(type: string, constructor: new (id: string) => Widget) {
        this.widgets[type] = constructor;
    }

    public getWidget(type: string, id: string): Widget {
        let constructor = this.widgets[type];
        return new constructor(id);
    }
}

let wm = new WidgetManager();

export var widgetManager: WidgetManager = wm;
