import * as CommerceTools from "../../sdk/nodejs/bin";

new CommerceTools.Subscription("my-test", {
    key: 'test2',
    destination: {
        type: "azure_servicebus",
        connectionString: "Endpoint=sb://unplatform8aaee3f5.servicebus.windows.net/;SharedAccessKeyName=test;SharedAccessKey=VE1/GVfCFSgUZYiMVo9laznJgZct+sp47Q7krhYpHR4=;EntityPath=test",
        region: "",
        uri: ""
    },
    messages: [{
        resourceTypeId: "order",
        types: ["OrderCreated"]
    }]
})

