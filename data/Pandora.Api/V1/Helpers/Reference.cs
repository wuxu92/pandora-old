namespace Pandora.Api.V1.Helpers;

public class Reference
{
    public static string? ObjectReference(object? input)
    {
        return input?.GetType().Name;
    }
}