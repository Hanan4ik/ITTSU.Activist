import { Tabs } from "expo-router";

import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import Fontisto from '@expo/vector-icons/Fontisto';
import Entypo from '@expo/vector-icons/Entypo';

export default function TabsLayout(){
    return <Tabs>
        <Tabs.Screen
        name='index'
        options={{
            title: "Мероприятия",
            tabBarIcon: ({focused}) => <MaterialIcons name="event-available" size={focused ? 28: 24} color="black" />,
            tabBarActiveBackgroundColor: '#dadadcff'
        }}
        />
        <Tabs.Screen
        name='profile'
        options={{
            title: "Профиль",
            tabBarIcon: ({focused}) => <Fontisto name="person" size={focused ? 28 : 24} color="black" />,
            tabBarActiveBackgroundColor: '#dadadcff',
        }}
        />
        <Tabs.Screen
        name='support'
        options={{
            title: "Задать вопрос",
            tabBarIcon: ({focused}) => <Entypo name="help" size={focused ? 28 : 24} color="black"/>,
            tabBarActiveBackgroundColor: '#dadadcff'
        }}
        />
    </Tabs>
}