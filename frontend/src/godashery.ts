import {log} from "./utils";
import {WidgetData, WidgetList, widgetManager} from "./widget";
import "./widgetloader";

export class GoDashery {
    private socket: WebSocket;
    private server: string;
    private reconnectInterval: number;
    private widgets: WidgetList = {};

    constructor() {
        this.reconnectInterval = 2500;
        this.parseServer();
        var id = "clock";
        this.widgets[id] = widgetManager.getWidget("ClockWidget", id);
    }

    public run() {
        log("Starting up GoDashery");
        this.connect();
    }

    private parseServer() {
        let host = window.location.host;
        let proto = (window.location.protocol === "https:" ? "wss" : "ws");

        let server = proto + "://" + host + "/data";

        log("WS server is at " + server);

        this.server = server;
    }

    private connect() {
        log("Connecting to server ", this.server);
        this.socket = new WebSocket(this.server);
        this.socket.onmessage = this.onMessage.bind(this);
        this.socket.onclose = this.onClose.bind(this);
    }

    private onMessage(event: MessageEvent) {
        let data = JSON.parse(event.data);

        if (data.type === "widgetData") {
            this.processWidgetData(data);
        } else {

            log("Event " + data.type);
            console.dir(data);

        }
    }

    private processWidgetData(data: WidgetData) {
        if (!this.widgets[data.id]) {
            log("Got event for widget " + data.id + " but couldn't find it.");
            return;
        }

        let widget = this.widgets[data.id];

        widget.update(data);
    }

    private reconnect() {
        log("Reconnecting...");
        this.connect();
    }

    private onClose(event: CloseEvent) {
        log("Disconnected from server", event);
        setTimeout(this.reconnect.bind(this), this.reconnectInterval);
    }
}