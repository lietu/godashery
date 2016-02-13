import {log} from "../utils";
import {Widget, WidgetData, BaseWidget, widgetManager} from "../widget";


interface ClockWidgetData extends WidgetData {
    time: string;
}

class ClockWidget extends BaseWidget implements Widget {
    constructor(id: string) {
        super(id);
    }

    public update(data: ClockWidgetData) {
        log("The time is " + data.time);
    }
}

widgetManager.registerWidget("ClockWidget", ClockWidget);
