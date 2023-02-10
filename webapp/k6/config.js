export const BaseUrl = "http://localhost:80";
export const Username = "asa";
export const Password = "asasasa";
export const TargetUsername = "mary"

export const getRandomAccount = () => {
  const accounts = ["mary", "jordan", "lula", "mai", "julia"];
  const name = accounts[Math.floor(Math.random() * accounts.length)];
  return [name, name + name];
}