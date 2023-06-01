// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Core.Exceptions.User;

public class UserLockedException : System.Exception
{
    public UserLockedException(string message) : base(message)
    {
    }
}
