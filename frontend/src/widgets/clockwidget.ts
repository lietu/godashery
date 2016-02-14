import blocks = require("blocks");
import {log} from "../utils";
import {Widget, WidgetData, BaseWidget, widgetManager} from "../widget";

interface TemplateData {
    time: any;
}

interface ClockWidgetData extends WidgetData {
    time: string;
}

class ClockWidget extends BaseWidget implements Widget {
    private templateData: TemplateData;

    constructor(id: string) {
        this.templateData = {
            "time": blocks.observable("00:00:00")
        };

        blocks.query(this.templateData);

        super(id);
    }

    public update(data: ClockWidgetData) {
        this.templateData.time.__value__ = data.time;
        this.templateData.time.update();
        log("The time is " + data.time);
    }
}

widgetManager.registerWidget("ClockWidget", ClockWidget);
