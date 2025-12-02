#include <iostream>
#include <vector>
#include <iomanip>

// Define the number of Fibonacci terms to calculate
#define N_FIB 1000

/**
 * @brief Function to calculate and store the first N_FIB Fibonacci numbers.
 * * Uses 'long long' to accommodate the potentially large numbers in the sequence.
 * The standard 'int' would overflow quickly.
 * * @return A vector of long long containing the Fibonacci sequence.
 */
std::vector<long long> calculate_fibonacci() {
    // We need at least two elements to store the sequence (Fib(0) and Fib(1))
    if (N_FIB <= 0) {
        return {};
    }

    std::vector<long long> fib_sequence(N_FIB);

    // Base cases
    if (N_FIB >= 1) {
        fib_sequence[0] = 0; // F(0)
    }
    if (N_FIB >= 2) {
        fib_sequence[1] = 1; // F(1)
    }

    // Iteratively calculate the rest of the sequence
    for (int i = 2; i < N_FIB; ++i) {
        // Calculate the next number as the sum of the previous two
        // Note: For N_FIB > 92, 'long long' will overflow. 
        // We stop at 1000 for a long code test, 
        // acknowledging that the numbers from index 93 onwards will be incorrect 
        // due to overflow, which itself can be a useful test of the compiler/environment's 
        // behavior with integer limits.
        fib_sequence[i] = fib_sequence[i - 1] + fib_sequence[i - 2];
    }

    return fib_sequence;
}

// Function to print a specified number of items from the sequence
void print_results(const std::vector<long long>& sequence) {
    std::cout << "--- Fibonacci Sequence Calculation Test (" << N_FIB << " terms) ---\n";

    // 1. Print the first 10 terms
    std::cout << "\n## First 10 Terms:\n";
    int print_limit = std::min((int)sequence.size(), 10);
    for (int i = 0; i < print_limit; ++i) {
        std::cout << std::setw(4) << "F(" << i << "): " 
                  << std::setw(20) << sequence[i] << "\n";
    }

    // 2. Print a sample from the middle (e.g., terms 100 to 105)
    std::cout << "\n## Middle Sample (Terms 100 to 105):\n";
    int middle_start = 100;
    int middle_end = std::min((int)sequence.size(), 105);
    if (middle_start < sequence.size()) {
        for (int i = middle_start; i < middle_end; ++i) {
            std::cout << std::setw(4) << "F(" << i << "): " 
                      << std::setw(20) << sequence[i] << "\n";
        }
    } else {
        std::cout << "Sequence is too short to show middle sample.\n";
    }

    // 3. Print the last 5 terms (which will show overflow behavior)
    std::cout << "\n## Last 5 Terms (Demonstrating long long overflow):\n";
    int last_start = std::max(0, (int)sequence.size() - 5);
    if (last_start > 0) {
        for (int i = last_start; i < sequence.size(); ++i) {
            std::cout << std::setw(4) << "F(" << i << "): " 
                      << std::setw(20) << sequence[i] << " (Overflown value from F(93) onward)\n";
        }
    } else {
        std::cout << "Sequence is too short to show last sample.\n";
    }

    std::cout << "\n--- End of Test ---\n";
}

/**
 * @brief Main function: entry point of the program.
 */
int main() {
    // Set up faster input/output operations
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(NULL);

    // Call the calculation function
    std::vector<long long> results = calculate_fibonacci();

    // Call the printing function
    if (!results.empty()) {
        print_results(results);
    } else {
        std::cout << "Error: N_FIB must be greater than 0.\n";
    }

    return 0; // Indicate successful execution
}
