
import { useState } from 'react';
import { Text, TextInput, View, StyleSheet, KeyboardAvoidingView, Platform} from "react-native";

export default function ProfilePage(){

    const [name, setName] = useState('');
    const [id, setId] = useState(0);
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    

    return ( 
        <Text>Smth</Text>
      );
}

const styles = StyleSheet.create(
    {
        baseView: {
            flex: 1,
            justifyContent: "center",
            alignItems: "center"
        },
        textInput: {
            height: 40,
            borderColor: 'blue',
            borderWidth: 2,
            borderRadius: 6
        }
    }
)