export const extractUserFromEmailAddress = (emailAddress) => {
    let atIndex = emailAddress.indexOf('@');
    if(atIndex === -1) {
        atIndex = emailAddress.length;
    }
    return emailAddress.substring(0, atIndex);
};
