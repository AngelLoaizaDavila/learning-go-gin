function changeMessage() {
  const messageElement = document.getElementById("message");
  const currentMessage = messageElement.textContent;

  const alternateMessages = [
    "Hello, world!",
    "¡Hola, mundo!",
    "Bonjour, le monde!",
    "Hallo, Welt!",
    "Ciao, mondo!",
    "こんにちは、世界！",
    "안녕하세요, 세계!",
    "Olá, mundo!",
    "Привет, мир!",
    "Hej, världen!",
  ];

  let newMessage;
  do {
    newMessage =
      alternateMessages[Math.floor(Math.random() * alternateMessages.length)];
  } while (newMessage === currentMessage);

  messageElement.textContent = newMessage;
}
