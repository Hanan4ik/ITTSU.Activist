import { Text, ScrollView } from "react-native";
import Event from "../types/event";
import { useState } from "react";

const API_URL = "http://127.0.0.1:8000"

export default function EventsPage() {
    const eventData = {
        images: [
            "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fantonenko-media.com%2Fwp-content%2Fuploads%2F2023%2F04%2Fphoto-event01845.jpg&f=1&nofb=1&ipt=6c19f157aecc95ac47209c76271e74ba172cd8f95cc7e634ea7996000f13191d",
            "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Feventum-premo.ru%2Fsites%2Fdefault%2Ffiles%2Fwiki%2F925e9c4cf9b7a9ee95214fb1df955efce2c80384.jpg&f=1&nofb=1&ipt=1d5afd0f9d781fb9c3dc0f517d785d23a6ab3274bd6b2f68bdea4fc2be4b4ba1"
        ],
        description: "Короче пиздато, круто приходите djhfasjhfsjdhfsdhfjkasdhfjksadhfjkh jksd hfjksdhfjk  hdfjka hsfkjahsd fhsjkd hjksdhf jkshdfjksh ajkfhaskj fdhajk shfk",
        averageRaiting: 4.5,

        
    }
    return (
        <ScrollView>
        <Event
        images={eventData.images}
        description={eventData.description}
        onRespond={() => console.log("Clicked on button")}
        averageRating={eventData.averageRaiting}/>
        </ScrollView>
    );
}
