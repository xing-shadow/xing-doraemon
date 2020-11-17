const TokenName = "Token"

export function SaveToken(token) {
    localStorage.setItem(TokenName,token)
}

export function GetToken() {
    const token = localStorage.getItem(TokenName)
    return token
}

export function DeleteToken() {
    localStorage.removeItem(TokenName)
}