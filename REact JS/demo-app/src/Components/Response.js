// response.js

const response = {
  hello: "Hey there! 👋 How can I help you today?",
  help: "I can chat about time, jokes, weather, math, motivation, food, queues, and more.",
  time: `Current time is ${new Date().toLocaleTimeString()}`,
  date: `Today is ${new Date().toDateString()}`,
  who: "I'm your offline AI sidekick 🤖. No Wi-Fi? Still got brains.",
  queue: "Queue system is healthy. Token retrieved, Captain.",
  weather:
    "I can't give real weather, but I'd guess it's sunny inside this app 🌤️",
  joke: "Why did the developer go broke? Because he used up all his cache. 😄",
  motivate: "Keep pushing—you're one clever bug away from brilliance 💡",
  food: "Tough choice... pizza, noodles, or plantains? You decide 🍕🍜🌱",
  math: "2 + 2? Easy: 4. Algebra? Try me.",
  default:
    "Hmm, I'm not sure how to respond to that yet this young man has not taught me well to reply that ... try saying hello or ask for help, time, weather, jokes, etc.",
};

export function getOfflineBotReply(message = "", ) {
  const msg = message.toLowerCase();

  // Keyword categories
  if (msg.includes("hello") || msg.includes("hi")) return response.hello;
  if (msg.includes("help")) return response.help;
  if (msg.includes("time")) return response.time;
  if (msg.includes("date")) return response.date;
  if (msg.includes("who")) return response.who;
  if (msg.includes("queue") || msg.includes("token")) return response.queue;
  if (msg.includes("weather")) return response.weather;
  if (msg.includes("joke") || msg.includes("funny")) return response.joke;
  if (msg.includes("motivate") || msg.includes("inspire"))
    return response.motivate;
  if (msg.includes("food") || msg.includes("hungry")) return response.food;
  if (msg.includes("math") || msg.includes("2 + 2") || msg.includes("add"))
    return response.math;

  // Fallback
  return response.default;
}
