// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using System.Reflection;
using CoproXpert.Core.Attributes;
using CoproXpert.Core.Enums;

namespace CoproXpert.Api.Sources.Helpers;

/// <summary>
///     Initializes the services
/// </summary>
public class ServiceInitializer
{
    private readonly List<string> _targetNamespaces = new() { "CoproXpert" };

    /// <summary>
    ///     Initializes the services in the specified service collection
    /// </summary>
    /// <param name="serviceCollection"></param>
    public void Init(IServiceCollection serviceCollection)
    {
        // Get all the services in the target namespaces and their subfolders
        var services = GetServicesInNamespaces(_targetNamespaces);

        // Loop through the services and add them to the service collection
        foreach (var service in services)
        {
            // Retrieve the service attribute
            var serviceAttribute = service.GetCustomAttribute<AutowireAttribute>();
            switch (serviceAttribute?.Lifetime)
            {
                case Lifetime.Scoped:
                    serviceCollection.AddScoped(service);
                    break;
                case Lifetime.Transient:
                    serviceCollection.AddTransient(service);
                    break;
                case Lifetime.Singleton:
                    serviceCollection.AddSingleton(service);
                    break;
                case null:
                    break;
            }
        }
    }

    private List<Type> GetServicesInNamespaces(ICollection<string> targetNamespaces)
    {
        // Get all the currently loaded assemblies
        var assemblies = AppDomain.CurrentDomain.GetAssemblies();

        // Get all the types in the assemblies that match any of the target namespaces and their subfolders,
        // and exclude the specified types
        var services = assemblies
            .SelectMany(assembly => GetTypesInNamespaces(assembly, targetNamespaces))
            .Where(type => type.GetCustomAttribute<AutowireAttribute>() != null)
            .ToList();

        return services;
    }

    private IEnumerable<Type> GetTypesInNamespaces(Assembly assembly, IEnumerable<string> targetNamespaces)
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
