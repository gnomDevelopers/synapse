//all validator functions
import { type IValidator } from "./constants";
import { FCapitalize } from "./functions";

//validate userLogin
export function ValidUserLogin(value: string): IValidator{
  if(value.match(/^[a-zA-Z0-9]+$/) === null){
    return {value: '', error: 'Incorrect login!'};
  }
  if(value.match(/[a-zA-Z]+/) === null){
    return {value: '', error: 'Login must contain one or more latin letters!'};
  }
  if(value.length < 4){
    return {value: '', error: 'Login is too short!'};
  }
  if(value.length > 25){
    return {value: '', error: 'Login is too long!'}; 
  }
  return {value: value, error: ''};
}

//validate userPassword
export function ValidUserPassword(value: string): IValidator{
  // if(value.match(/^[a-zA-Z0-9-]+$/) === null){
  //   return {value: '', error: 'Incorrect password!'};
  // }
  // if(value.match(/[a-zA-Z]+/) === null){
  //   return {value: '', error: 'Password must contain latin letters!'};
  // }
  // if(value.match(/[a-z]+/) === null){
  //   return {value: '', error: 'Password must contain a lowercase letter!'};
  // }
  // if(value.match(/[A-Z]+/) === null){
  //   return {value: '', error: 'Password must contain an uppercase letter!'};
  // }
  // if(value.match(/[0-9]+/) === null){
  //   return {value: '', error: 'Password must contain a number!'};
  // }
  // if(value.length < 8){
  if(value.length < 4){
    return {value: '', error: 'Password is too short!'};
  }
  if(value.length > 50){
    return {value: '', error: 'Password is too long!'};
  }
  return {value: value, error: ''};
}

//validate userEmail
export function ValidUserEmail(value: string): IValidator{
  if(value.toLowerCase().match(/^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/) === null){
    return {value: '', error: 'Incorrect email address!'};
  }
  return {value: value, error: ''}
}

//validate userName
export function ValidUserName(value: string): IValidator{
  if(value.match(/^[а-яА-Яa-zA-ZёЁ]+([-'`]{1}[а-яА-Яa-zA-ZёЁ]+)?$/) === null){
    return {value: '', error: 'Incorrect name!'};
  }
  if(value.length < 2){
    return {value: '', error: 'Your name is too short!'};
  }
  if(value.length > 30){
    return {value: '', error: 'Your name is too long!'};
  }
  return {value: FCapitalize(value.toLowerCase()), error: ''}
}

//validate userSurname
export function ValidUserSurname(value: string): IValidator{
  if(value.match(/^[а-яА-Яa-zA-ZёЁ]+([-'`]{1}[а-яА-Яa-zA-ZёЁ]+)?$/) === null){
    return {value: '', error: 'Incorrect surname!'};
  }
  if(value.length < 2){
    return {value: '', error: 'Your surname is too short!'};
  }
  if(value.length > 30){
    return {value: '', error: 'Your surname is too long!'};
  }
  return {value: FCapitalize(value.toLowerCase()), error: ''}
}

//validate userThirdname
export function ValidUserThirdname(value: string): IValidator{
  if(value.length === 0){
    return {value: '', error: ''}
  }
  if(value.match(/^[а-яА-Яa-zA-ZёЁ]+([-'`]{1}[а-яА-Яa-zA-ZёЁ]+)?$/) === null){
    return {value: '', error: 'Incorrect patronymic!'};
  }
  if(value.length < 2){
    return {value: '', error: 'Your patronymic is too short!'};
  }
  if(value.length > 30){
    return {value: '', error: 'Your patronymic is too long!'};
    
  }
  return {value: FCapitalize(value.toLowerCase()), error: ''}
}

export function ValidUserFIO(value: string): IValidator{
  const splitted = value.split(' ')
  if(splitted.length !== 3) return {value: '', error: 'Incorrect fio!'};
  const [name, surname, thirdname] = splitted;

  if (typeof name === "undefined") {
    return {value: '', error: "Name missed!"}
  }
  if (typeof surname === "undefined") {
    return {value: '', error: "Surname missed!"}
  }
  if (typeof thirdname === "undefined") {
    return {value: '', error: "Thirdname missed!"}
  }

  const nameValid = ValidUserName(name);
  const surnameValid = ValidUserSurname(surname);
  const thirdnameValid = ValidUserThirdname(thirdname);

  if(nameValid.error === '' && surnameValid.error === '' && thirdnameValid.error === ''){
    return {value: value, error: ''};
  }
  else{
    return {value: '', error: 'Incorrect fio!'};
  }
}

export function ValidUserTelegram(value: string): IValidator{
  if (value.length < 5 || value.length > 32 || !(/^[a-zA-Z0-9_]+$/.test(value)) ) {
    return {value: '', error: 'Incorrect telegram tag!'};
  }
  return {value: value, error: ''};
}