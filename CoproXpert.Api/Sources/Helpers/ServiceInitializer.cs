// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Reflection;
using CoproXpert.Database.Models;
using CoproXpert.Database.Repositories;

namespace CoproXpert.Api.Sources.Helpers;

/// <summary>
///     Initializes the services
/// </summary>
public static class ServiceInitializer
{
    private static readonly List<string> s_targetNamespaces = new() { "CoproXpert.Database.Repositories" };

    private static readonly List<Type> s_excludedTypes = new() { typeof(BaseModel), typeof(BaseRepository<>) };

    /// <summary>
    ///     Initializes the services in the specified service collection
    /// </summary>
    /// <param name="serviceCollection"></param>
    public static void Init(IServiceCollection serviceCollection)
    {
        // Get all the services in the target namespaces and their subfolders
        var services = GetServicesInNamespaces(s_targetNamespaces, s_excludedTypes);

        // Loop through the services and add them to the service collection
        foreach (var service in services)
        {
            serviceCollection.AddScoped(service);
        }
    }

    private static List<Type> GetServicesInNamespaces(ICollection<string> targetNamespaces,
        ICollection<Type> excludedTypes)
    {
        // Get all the currently loaded assemblies
        var assemblies = AppDomain.CurrentDomain.GetAssemblies();

        // Get all the types in the assemblies that match any of the target namespaces and their subfolders,
        // and exclude the specified types
        var services = assemblies
            .SelectMany(assembly => GetTypesInNamespaces(assembly, targetNamespaces))
            .Where(type => !excludedTypes.Contains(type))
            .ToList();

        return services;
    }

    private static IEnumerable<Type> GetTypesInNamespaces(Assembly assembly, ICollection<string> targetNamespaces)
    {
        // Get all the types in the assembly
        var assemblyTypes = assembly.GetTypes();

        // Loop through the target namespaces and their subfolders
        foreach (var targetNamespace in targetNamespaces)
        {
            // Get all the types in the assembly that match the target namespace and its subfolders
            var types = assemblyTypes
                .Where(type => type.Namespace != null && type.Namespace.StartsWith(targetNamespace))
                .ToList();

            foreach (var type in types)
            {
                yield return type;
            }
        }
    }
}
