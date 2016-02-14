import blocks = require("blocks");
import {log} from "../utils";
import {Widget, WidgetData, BaseWidget, widgetManager} from "../widget";

var templateData = {
    "time": "00:00:00"
};

blocks.query(templateData);

console.dir(blocks);

interface ClockWidgetData extends WidgetData {
    time: string;
}

class ClockWidget extends BaseWidget implements Widget {
    constructor(id: string) {
        super(id);
    }

    public update(data: ClockWidgetData) {
        templateData.time = data.time;
        log("The time is " + data.time);
    }
}

widgetManager.registerWidget("ClockWidget", ClockWidget);
