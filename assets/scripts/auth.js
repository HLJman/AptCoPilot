

var params = {
    region: "eu-west-2",
    userPoolID: "us-west-2_Sxqf852KB",
    clientID: "7i12nnhnirp860t1o0oju4run1"
};

function login(username, password, callback) {
    var authenticationData = {
        Username: username,
        Password: password
    };

    var authenticationDetails = new AWSCognito.CognitoIdentityServiceProvider.AuthenticationDetails(authenticationData);
    var poolData = {
        UserPoolId: params.userPoolID,
        ClientId: params.clientID
    };
    var userPool = new AWSCognito.CognitoIdentityServiceProvider.CognitoUserPool(poolData);
    var userData = {
        Username: username,
        Pool: userPool
    };
    var cognitoUser = new AWSCognito.CognitoIdentityServiceProvider.CognitoUser(userData);
    cognitoUser.authenticateUser(authenticationDetails, {
        onSuccess: function (result) {
            console.log(result)
            console.log('access token + ' + result.getAccessToken().getJwtToken());
            /*Use the idToken for Logins Map when Federating User Pools with Cognito Identity or when passing through an Authorization Header to an API Gateway Authorizer*/
            console.log('idToken + ' + result.idToken.jwtToken);

            callback(result)
        },

        onFailure: function (err) {
            alert(err);
        },

    });
}

function signup() {
    var attributeList = [];

    var dataEmail = {
        Name: 'email',
        Value: '...' // your email here
    };
    var dataPhoneNumber = {
        Name: 'phone_number',
        Value: '...' // your phone number here with +country code and no delimiters in front
    };
    var attributeEmail =
        new AWSCognito.CognitoIdentityServiceProvider.CognitoUserAttribute(dataEmail);
    var attributePhoneNumber =
        new AWSCognito.CognitoIdentityServiceProvider.CognitoUserAttribute(dataPhoneNumber);

    attributeList.push(attributeEmail);
    attributeList.push(attributePhoneNumber);

    var cognitoUser;
    userPool.signUp('username', 'password', attributeList, null, function (err, result) {
        if (err) {
            alert(err);
            return;
        }
        cognitoUser = result.user;
        console.log('user name is ' + cognitoUser.getUsername());
    });
}

function confirmSignup(code) {
    cognitoUser.confirmRegistration(code, true, function (err, result) {
        if (err) {
            alert(err);
            return;
        }
        console.log('call result: ' + result);
    });
}