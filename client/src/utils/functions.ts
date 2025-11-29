export function FCapitalize(str: string): string{
  if (str.length < 2) return str.toUpperCase();
  return str[0]!.toUpperCase() + str.slice(1);
}

export function FormatUTCDateStringToLocal(utcDateString: string): string { //2025-04-13T17:40:01.028413Z -> 13.04.2025 20:40
  const date = new Date(utcDateString);
  const day = String(date.getDate()).padStart(2, '0');
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const year = date.getFullYear();
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  return `${day}.${month}.${year} ${hours}:${minutes}`;
}

export function CHECK_COOKIE_EXIST(cookieName: string): boolean{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    if(cookie.split('=')[0]!.trim() === cookieName) return true;
  }
  return false;
}

export function GET_COOKIE(cookieName: string): string{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    let [key, value] = cookie.split('=');
    if(key!.trim() === cookieName) return value!.trim();
  }
  return '';
}

export function SET_COOKIE(name: string, value: string, expires: Date) {
  let cookieName = encodeURIComponent(name) + "=" + encodeURIComponent(value);
  document.cookie = `${cookieName}; expires=${expires.toUTCString()}; samesite=lax; path=/`;
}