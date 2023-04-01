using CoproXpert.Database.Models;
using CoProXpert.Database.Models;
using CoproXpert.Database.Services;

namespace CoproXpert.Api.Sources.Helpers;

public static class ServiceInitializer
{
    private static readonly List<string> TargetNamespaces = new() { "CoproXpert.Database.Services" };

    private static readonly List<Type> ExcludedTypes = new()
        { typeof(BaseModel), typeof(BaseService<>), typeof(IModel) };

    public static void Init(IServiceCollection serviceCollection)
    {
        // Get all the services in the target namespace
        var services = GetServicesInNamespace(TargetNamespaces, ExcludedTypes);

        // Loop through the services and add them to the service collection
        foreach (var service in services) serviceCollection.AddScoped(service);
    }

    private static List<Type> GetServicesInNamespace(ICollection<string> targetNamespaces,
        ICollection<Type> excludedTypes)
    {
        // Get all the currently loaded assemblies
        var assemblies = AppDomain.CurrentDomain.GetAssemblies();

        // Get all the types in the assemblies that match any of the target namespaces and exclude the specified types
        var services = assemblies
            .SelectMany(assembly => assembly.GetTypes())
            .Where(type => type.Namespace != null
                           && targetNamespaces.Any(targetNamespace => type.Namespace.Equals(targetNamespace))
                           && !excludedTypes.Contains(type))
            .ToList();

        return services;
    }
}