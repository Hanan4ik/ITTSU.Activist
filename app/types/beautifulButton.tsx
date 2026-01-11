import { TouchableOpacity, Text, StyleSheet } from "react-native";

/**
 * 
 * @param btnProps map with fields
 * color: clr
 * onPress: func
 * @param btnTextProps map with fields
 * color: clr
 * fontSize: number
 * text: string
 * @returns Beautiful button
 */
const BeautifulButton = ({btnProps, btnTextProps}) => {

    const styles = StyleSheet.create({
        buttonText: {
            color: btnTextProps.color,
            fontSize: btnTextProps.fontSize,
            fontWeight: 'bold',
        },
        button: {
            backgroundColor: btnProps.color,
            paddingVertical: 15,
            borderRadius: 8,
            alignItems: 'center',
            marginTop: 10
        },
        });
    
    return (
        <TouchableOpacity style={styles.button} onPress={btnProps.onPress}>
            <Text style={styles.buttonText}>{btnTextProps.text}</Text>
        </TouchableOpacity>
    );
}

export default BeautifulButton;